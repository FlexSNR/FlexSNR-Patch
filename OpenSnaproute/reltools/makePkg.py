import os
import sys
import json
import time
import fileinput
import subprocess
import getpass
import argparse
from fabric.api import env,local,run,parallel
from optparse import OptionParser
from curator.personality import FlexPersonality


PACKAGE_BUILD = "PKG_BUILD=TRUE"
TEMPLATE_BUILD_TYPE = "PKG_BUILD=FALSE"
TEMPLATE_CHANGELOG_VER = "0.0.1"
TEMPLATE_BUILD_DIR = "flexswitch-0.0.1"
TEMPLATE_BUILD_TARGET = "cel_redstone"
TEMPLATE_PLATFORM_BUILD_TARGET = "dummy"
TEMPLATE_ALL_TARGET = "ALL_DEPS=buildinfogen codegen installdir ipc exe install"
PKG_ONLY_ALL_TARGET = "ALL_DEPS=installdir install"

platformHandlers = {
    'ingrasys_s9100'    : FlexPersonality()
}

def buildDocker(command):
    p = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    out, err = p.communicate()
    print out.rstrip(), err.rstrip()
    print "Docker image return code ", p.returncode
    print "Check version of  image -- docker images"
    return

def executeCommand(command):
    out = ''
    if type(command) != list:
        command = [command]
    for cmd in command:
        print 'Executing command %s' %(cmd)
        local(cmd)
    return out

if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='FlexSwitch Package builder')
    parser.add_argument('-r', '--release',
                        dest='release',
                        action='store_true',
                        default=False,
                        help='Is Release')
    parser.add_argument('--platform',
                        type=str,
                        dest='platform',
                        action='store',
                        nargs='?',
                        default="",
                        help='device platform')


    args = parser.parse_args()
    if args.release == True:
        usrName = 'release'
    else:
        usrName = getpass.getuser()

    if args.platform != '':
        buildPlatform = args.platform
    else:
        buildPlatform = ''

    with open("pkgInfo.json", "r") as cfgFile:
        pkgInfo = cfgFile.read().replace('\n', '')
        parsedPkgInfo = json.loads(pkgInfo)
    cfgFile.close()
    firstBuild = True
    buildTargetList = parsedPkgInfo['platforms']
    pkgVersion = usrName + '_' + parsedPkgInfo['major']+ '.'\
                  + parsedPkgInfo['minor'] +  '.' + parsedPkgInfo['patch'] + \
                  '.' + parsedPkgInfo['build'] + '.' + parsedPkgInfo['changeindex']
    pkgVersionNum = parsedPkgInfo['major']+ '.'\
                  + parsedPkgInfo['minor'] +  '.' + parsedPkgInfo['patch'] + \
                  '.' + parsedPkgInfo['build'] + '.' + parsedPkgInfo['changeindex']
    build_dir = "flexswitch-" + pkgVersion
    command = [
            'rm -rf ' + build_dir,
            'make clean_all'
            ]
    executeCommand(command)
    startTime = time.time()
    for buildTargetDetail in buildTargetList:
        buildTarget = buildTargetDetail['odm']
        if buildTarget == buildPlatform or buildPlatform == '':
            platform = buildTargetDetail['platform']
            print "Building pkg for Tgt:%s Platform %s" %(buildTarget, platform)
            platfomHdlr = platformHandlers.get(buildTarget, None)
            pkgName = "flexswitch_" + buildTarget + "-" + pkgVersion + "_amd64.deb"
            if firstBuild:
                preProcess = [
                        'cp -a tmplPkgDir ' + build_dir,
                        'cp Makefile ' + build_dir,
                        'sed -i s/' + TEMPLATE_BUILD_DIR +'/' + build_dir + '/ ' + build_dir +'/Makefile',
                        'sed -i s/' + TEMPLATE_BUILD_TYPE +'/' + PACKAGE_BUILD + '/ ' + build_dir + '/Makefile',
                        'sed -i s/' + TEMPLATE_CHANGELOG_VER +'/' + pkgVersionNum+ '/ ' + build_dir + '/debian/changelog',
                        'sed -i s/' + TEMPLATE_BUILD_TARGET +'/' + buildTarget + '/ ' + build_dir + '/Makefile',
                        'sed -i s/' + TEMPLATE_PLATFORM_BUILD_TARGET +'/' + platform + '/ ' + build_dir + '/Makefile'
                        ]
                executeCommand(preProcess)
                #Build all binaries only once
                os.chdir(build_dir)
                executeCommand('make all')
                os.chdir("..")
                executeCommand('python buildInfoGen.py')
                firstBuild = False
                #Change all target prereqs
                for line in fileinput.input(build_dir+'/Makefile', inplace=1):
                    print line.replace(TEMPLATE_ALL_TARGET, PKG_ONLY_ALL_TARGET).rstrip('\n')
            else:
                #Change build target and all target prereqs
                preProcess = [
                        'sed -i s/' + prevBldTgt +'/' + buildTarget + '/ ' + build_dir + '/Makefile',
                        'sed -i s/' + prevPlatTgt +'/' + platform + '/ ' + build_dir + '/Makefile'
                        ]
                executeCommand(preProcess)
                os.chdir(build_dir)
                executeCommand('make asicd')
                os.chdir("..")
            prevBldTgt = buildTarget
            prevPlatTgt = platform
            os.chdir(build_dir)
            pkgRecipe = [
                    'fakeroot debian/rules clean',
                    'fakeroot debian/rules build',
                    ]
            executeCommand(pkgRecipe)
            if platfomHdlr:
                platfomHdlr.performBuildTimeCustomization(build_dir)
            executeCommand('fakeroot debian/rules binary')
            os.chdir("..")
            cmd = 'mv flexswitch_' + pkgVersionNum + '*_amd64.deb ' + pkgName
            local(cmd)
            if buildTarget == "docker":
                 cmd = 'python dockerGen/buildDocker.py'
                 print "Building Docker image with flex package ", pkgName
                 buildDocker(cmd + " " + pkgName)
    command = [
            'rm -rf ' + build_dir,
            'make clean_all'
                ]
    executeCommand(command)
