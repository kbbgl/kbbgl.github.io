---
slug: macbook-pro-2020-high-cpu-caused-siri
title: MacBook Pro 2020 High CPU caused by Siri
description: Fill me up!
authors: [kbbgl]
tags: [apple,corespeechd,cpu,debugging,macos]
---

## Introduction

A few months ago, I received a highly-anticipated 2020 32GB, 2.3 GHz Quad-Core Intel Core i7 MacBook Pro. Highly-anticipated because I already had one stolen (a 2019 version) earlier last year in a robbery in an AirBnB apartment I was renting while I was staying in Barcelona. It was quite a dramatic story but I won’t get into the details. This is a tech blog after all.

Back to the new computer I acquired, the thing was flying. I could open large projects in IntelliJ and VSCode, run a Kubernetes cluster using [`minikube`](https://minikube.sigs.k8s.io/docs/start/) and endless amounts of terminals simultaneously on different workspaces and the thing would not miss a beat.

Within the first week of receiving machine, I got a notification indicating that there was a system update. I never had any fears come with system updates. I always found myself jumping ship and hoping no regressions or new issues are introduced upon upgrades/updates. I’m not a skeptic, I believe Apple’s QA is up to the highest testing standards.

Anyway, I let the update run for a few hours and after a couple of restarts, the OS was back up and could work on the machine again.

Pretty soon thereafter, I heard the machine fan working intensively and the low-carbon aluminum enclosure was excessively hot. I also noticed that I the interaction with the OS and applications was crawlingly-slow. I even had the machine crash 2-3 times! I decided I had enough and that it was time to put on my Sherlock Holmes hat and start the investigation.

## Finding the Culprit Process

It makes sense that the first place to review would be Activity Monitor. It allows us to check the amount of energy each running process is utilizing in addition to CPU cycles and RAM allocation. Below is the screenshot of what I saw after launching Activity Monitor:

![activity-monitor](https://tilsupport.files.wordpress.com/2021/03/image.png)

We can see that the top process, `ZscalerTunnel` was really using up the CPU. I knew what this process belonged to (a VPN service). I was able to fix the `ZscalerTunnel` CPU hog by installing the latest version of the application and restarting the machine.

But the second process, `corespeechd` was not letting go and was at times using between 3 to 4 CPUs (300-400%). And the worst part was that I had no idea what `corespeechd` was (aside from a guess that it was some sort of daemon because of the letter ‘d’ in the suffix).

## What is `corespeechd`?

This question was hard to find online. I believe because I am not a registered Apple developer, I could not see any relevant documentation about this process/service. What I did find was that many people had a problem with this daemon hogging machine resources such as [massive network utilization (1)](https://discussions.apple.com/thread/8643914?page=2) and [massive network utilization (2)](https://discussions.apple.com/thread/250955260?page=2). And I also found a lot of people experiencing the [high CPU consumption](https://forums.macrumors.com/threads/cpu-usage-corespeechd.2158710/) and on [Twitter as well](https://twitter.com/alecmuffett/status/1089721018015539200?lang=en).

According to the online research, all fingers were pointing to one feature: Siri. In addition, all recommendations to mitigate this issue mentioned disabling/enabling or turning off Siri completely. Unfortunately for me, corespeechd was still causing problems after my attempt to disable Siri.
Since the machine had crashed a few times, I decided the next step would be to review the system crash report.
What I found was that the OS crashed because of a segmentation fault:

```segfault
Crashed Thread:        0  Dispatch queue: com.apple.main-thread

Exception Type:        EXC_BAD_ACCESS (SIGSEGV)
Exception Codes:       KERN_INVALID_ADDRESS at 0x0000000108745050
Exception Note:        EXC_CORPSE_NOTIFY

Termination Signal:    Segmentation fault: 11
Termination Reason:    Namespace SIGNAL, Code 0xb
Terminating Process:   exc handler [11719]

Thread 0 Crashed:: Dispatch queue: com.apple.main-thread
0   libobjc.A.dylib                0x00007fff202639af objc_release + 15
1   com.apple.CoreFoundation       0x00007fff20478a76 -[__NSDictionaryI dealloc] + 146
2   libobjc.A.dylib                0x00007fff2028139d AutoreleasePoolPage::releaseUntil(objc_object**) + 167
3   libobjc.A.dylib                0x00007fff2026433e objc_autoreleasePoolPop + 161
4   com.apple.CoreFoundation       0x00007fff2047e1f0 _CFAutoreleasePoolPop + 22
5   com.apple.CoreFoundation       0x00007fff20587748 __CFRunLoopPerCalloutARPEnd + 41
6   com.apple.CoreFoundation       0x00007fff204bc88b __CFRunLoopRun + 2788
7   com.apple.CoreFoundation       0x00007fff204bb6ce CFRunLoopRunSpecific + 563
8   com.apple.Foundation           0x00007fff21248fa1 -[NSRunLoop(NSRunLoop) runMode:beforeDate:] + 212
9   com.apple.Foundation           0x00007fff212d7384 -[NSRunLoop(NSRunLoop) run] + 76
10  com.apple.authorizationhost    0x0000000108661727 main + 302
11  libdyld.dylib                  0x00007fff203e0621 start + 1
```

What was immediately visible was that there were instructions in memory address `0x00007fff20478a76` that were being allocated by the `com.apple.CoreFoundation` package. This gave me the evidence I needed that tied the system crash to the issues witnessed by the `Core` services.

I decided to run `corespeechd` using `strace` so I could see what system calls the process was executing during run-time. I found the following output indicating there was some permission issues attempting to access a certain process probe:

```text
dtrace: error on enabled probe ID 2378 (ID 918: syscall::kevent_id:return): invalid user access in action #5 at DIF offset 0
dtrace: error on enabled probe ID 2385 (ID 904: syscall::workq_kernreturn:return): invalid user access in action #5 at DIF offset 0
```

## Using `launchctl` to Disable `corespeechd`

Since Siri was already off, I was pretty sure that if I found a way to completely disable `corespeechd`, I would be able to release the daemon from eating up my processors.

I am familiar with using `systemctl` and `init.d` on Linux to be able to control system services. On Windows I usually just use the `Stop-Service` PowerShell cmdlet. But for Mac I wasn’t sure what command line tool needs to be used. I found out that the preferred way is to use `launchctl`.

I also wanted to understand where MacOS startup scripts are stored. I found out that there were 5 different directories:

- `/System/Library/LaunchDaemons/` – System-wide daemons provided by the operating system.
- `/System/Library/LaunchAgents/` – Per-user agents provided by the operating system.
- `~/Library/LaunchAgents/` – Per-user agents provided by the user.
- `/Library/LaunchAgents/` – Per-user agents provided by the administrator.
- `/Library/LaunchDaemons/` – System-wide daemons provided by the administrator.

`launchd` manages the processes, both for the system as a whole and for individual users using configuration files with the `.plist` extension.

I reviewed the `launchctl` `man` page and found that we can disable a service by using the `unload` command, for example:

```bash
sudo launchctl unload -w com.apple.corespeechd
```

That looked great! The `corespeechd` daemon was no longer running and my CPU was back to normal consumption. Unfortunately, I realized my celebrations were premature when I restarted my machine. `corespeechd` was back up and feasting on my machine resources again.

I went back to the `launchctl man` page and found that there was another option, `remove`, that could do the trick as it seemed to perform a persistent disabling of services (unlike the `unload` operation which was temporary until system restart relaunched all daemons). The problem was that I was unable to run the command:

```bash
sudo launchctl remove com.apple.corespeechd
```

because the operating system had [System Integrity Protection (SIP)](https://support.apple.com/en-us/HT204899) enabled. I found that I could use the csrutil command-line tool to interact with SIP. To interact with SIP, we need to go into Recovery Mode.

tl;dr ,these are the steps I took to completely disable `corespeechd`:

1. Reboot the machine.
1. Press and hold `Command + R`.
1. Once the Recovery menu loads up, in the top menu bar, select `Utilities > Terminal`.
1. Run the following command to disable System Integrity Protection (SIP):

```bash
csrutil disable
```

1. Reboot the machine.
1. Press and hold Command + R.
1. Once the Recovery menu loads up, in the top menu bar, select Utilities > Terminal.
1. Run the following command to disable System Integrity Protection (SIP):

You should see the following message to acknowledge that SIP was actually disabled:

```bash
System Integrity Protection status: disabled.
```

1. Reboot the machine.
1. When OS loads, open a Terminal and run the following command to disable `corespeechd` daemon:

```bash
sudo launchctl remove com.apple.corespeechd
```

You should see that corespeechd daemon is no longer running and consuming massive CPU!
