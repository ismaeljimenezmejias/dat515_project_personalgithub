# Lab 1: Unix Basics and Command Line Tools

| Lab 1:           | Unix Basics and Command Line Tools |
| ---------------- | ---------------------------------- |
| Subject:         | DAT515 Cloud Computing             |
| Deadline:        | **August 22, 2025 23:59**          |
| Expected effort: | 2-3 hours                          |
| Grading:         | Pass/fail                          |
| Submission:      | Individually                       |

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Get Access to OpenStack VM](#get-access-to-openstack-vm)
  - [Task: Add Your SSH Public Key to Your GitHub Repository](#task-add-your-ssh-public-key-to-your-github-repository)
  - [Task: Get OpenStack VM](#task-get-openstack-vm)
- [The Unix System at UiS](#the-unix-system-at-uis)
  - [Task: Sign up for Unix Account](#task-sign-up-for-unix-account)
- [The Missing Semester of Your CS Education by MIT](#the-missing-semester-of-your-cs-education-by-mit)
  - [Additional Resources and Tips](#additional-resources-and-tips)
  - [Task: Unix/Linux and Git Multiple-Choice Questions](#task-unixlinux-and-git-multiple-choice-questions)
- [Remote Login with Secure SHell: Unix](#remote-login-with-secure-shell-unix)
- [Configuring SSH for Passwordless Login](#configuring-ssh-for-passwordless-login)
  - [Task: Setting Up SSH Keys for Passwordless GitHub Authentication](#task-setting-up-ssh-keys-for-passwordless-github-authentication)
- [Cloning Your Assignments Repository and Submitting to QuickFeed](#cloning-your-assignments-repository-and-submitting-to-quickfeed)

## Introduction

Most lab assignments can be performed on your local machine (this is preferred).
If you already run Linux or macOS on your laptop, you should be ready to go.
Linux and macOS are to a large extent relatively similar at the command level.
If you are running Windows, please consult these [instructions](https://github.com/dat515-2025/info/blob/main/setup-wsl.md).

We do recommend a somewhat powerful machine with at least 16 GB of RAM.
If your machine does not meet these requirements, you may want to use the OpenStack VM provided for the course.

In the first part of this lab, we will get started with some of the Unix tools that will be used during all labs in this course.
Please read carefully the instructions provided in this part before starting to solve the lab assignments.

## Get Access to OpenStack VM

Some lab assignments will require you to use an OpenStack VM.
You may also use the OpenStack VM for other assignments, e.g., if your local machine is not powerful enough.
However, we do recommend that you prefer your local machine, since it will provide more flexibility.

### Task: Add Your SSH Public Key to Your GitHub Repository

To enable secure access to your course resources, you are required to add your SSH public key to your personal GitHub assignments repository.
Generate an SSH key pair (if you haven’t already) on your local machine:

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

This will create a public key file: `~/.ssh/id_ed25519.pub`.
You can inspect the contents of the public key file using the following command:

```sh
$ cat ~/.ssh/id_ed25519.pub
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEqjMwVp9QGB8+V1/BFt1XvhFXOIDepLyW+ZpxKbohzk meling@Heins-Mac-Studio.local
```

Next, copy your public key file into the `1unix` directory, naming it as `github-username.pub`.
Then commit and push the public key to your GitHub repository.

```bash
cd 1unix
# Replace github-username with your actual GitHub username
cp ~/.ssh/id_ed25519.pub github-username.pub
git add github-username.pub
git commit -m "added SSH public key for OpenStack VM access"
git push
```

Important:
Do NOT upload your private key (the file without .pub).
Only upload your public key.

### Task: Get OpenStack VM

Once you have signed up to QuickFeed, and completed the task above, we will provision an OpenStack VM for you.
We will use your GitHub username as your username for the VM.
Once we have provisioned the VM, you can connect to this VM to work on your assignments, as follows:

```bash
ssh github-username@152.94.171.171 -i ~/.ssh/id_ed25519
```

You do not need a passphrase; the command above is using the SSH private key corresponding to the public key you uploaded in the previous task to access your VM.

## The Unix System at UiS

We also recommend that you obtain a Unix account, as this will allow you to access various Unix servers that you may need throughout your studies in other courses.
However, it is not strictly necessary for this course.

### Task: Sign up for Unix Account

Get an account for Unix system by following these [instructions](http://user.ux.uis.no).
Be sure to read the section **Using the UNIX system**.

_It may take a while before you get access, but you can continue learning while you wait._

## The Missing Semester of Your CS Education by MIT

Throughout this and other courses and as a software engineer, you will often need to use command-line tools to interact with computers.
Lack of knowledge of the available tools will lead to manually performing repetitive tasks or spending lots of time searching online for solutions.
For these reasons and more, we suggest that you to go through [The Missing Semester of Your CS Education](https://missing.csail.mit.edu/) from MIT (hereafter referred to as the Missing Semester).
You can read more about the [motivation](https://missing.csail.mit.edu/about/) behind the course.

Fun fact: One of the lecturers in this course is a Norwegian, [Jon Gjengset](https://thesquareplanet.com).
He was a PhD student at MIT.
He is a YouTuber specializing in lengthy live-coding sessions about building stuff using the Rust programming language.

You should try to answer or at least understand the answers to the **Exercises** section at the end of each lecture.
Additionally, we give a set of multiple choice questions below, which mostly correspond to lectures 1, 2, 4, 5, and 6.

### Additional Resources and Tips

- [UNIX Tutorial for Beginners](http://www.ee.surrey.ac.uk/Teaching/Unix/).
  Eight simple tutorials covering the basics of various Unix/Linux commands.
  You may use these as a reference if you struggle to answer some of the questions or want a more in-depth overview than that offered by the Missing Semester.

- [Unix/Linux Command Reference](https://files.fosswire.com/2007/08/fwunixref.pdf).
  A cheat sheet of several frequently used Unix/Linux commands.

- Remember that almost every Unix/Linux command has a manual page, or man page for short, which can be accessed with `man` command, e.g. `man ls` for the `ls` command.

- _Tip:_ Use the `git help` command whenever in doubt about a Git command.
  It lets you read more about the functionality of each of Git's subcommands, e.g. by running `git help commit` for information about `git commit`, such as options, or `git help pull` for information about `git pull`.

- _Tip:_ Navigating `man`, `less` and `git help` buffers: The buffers opened by the `man`, `less` and `git help` commands support vi(m)-like navigation.
  - You can move down by one line by pressing the `Down` arrow key or the `j` key, or up by one line by pressing the `Up` arrow key or the `k` key.
  - You can move up or down by one page by pressing the `PageUp` and `PageDown` keys.
    Alternatively you can press the `f` ("forward") or `b` ("backward") keys.
  - You can go to the start or end of the buffer by pressing the respective `Home` and `End` keys.
    Alternatively you can press the `g` or `G` keys for the same functionality.
    There are often examples at the end of man pages.
  - You can search for some text by pressing the `/` key.
    Press `n` to go to the next match, and `N` to go to the previous match.

### Task: Unix/Linux and Git Multiple-Choice Questions

Answer the questions inline in the markdown files, as explained in the heading of each file.

1. [Questions related to the Missing Semester](./missing_semester_questions.md) lectures 1, 2 and 4.
2. [Shell questions](shell_questions.md).
   Some of these commands may not be covered by the Missing Semester lectures.
   We recommend reading the relevant man pages and checking the other related resources mentioned above.
3. [Questions about Git](./git_questions.md) based on lecture 6 of the Missing Semester as well as some regularly used Git commands.
   _Hint:_ Some of the questions may be heavily influenced by StackOverflow questions.

Note that, some commands behave differently on macOS and Linux, because they are based on different heritage.
Typically, macOS and Linux may sometimes use different flags for altering the behavior of a command.
We have made notes on these differences, where we are aware of them, but should you discover an incompatibility in these labs, please let us know.

Further, this lab was designed with the `bash` Unix shell, which is the default on Linux.
The default shell is `zsh` on macOS.
If you experience any issues related to running a different shell than `bash`, please try the same on Linux, and let us know.
To determine your shell, use the following command:

```console
echo $SHELL
```

## Remote Login with Secure SHell: Unix

_Skip this part if you haven't got a Unix account password yet._

To access UiS Unix machines, you should review the [UiS SSH information](https://foswiki.ux.uis.no/bin/view/Info/SshCommand).
Here is a [SSH tutorial video](https://youtu.be/qik3HHZW6C0) illustrating the steps below (and a bit more).

You can use `ssh` to log on to another machine without physically going to that machine and login there.
To log onto a machine using ssh, open a terminal window and type a command according to this template, and make sure to replace `username` and `hostname`:

```console
ssh username@hostname
```

If you are not on campus, e.g., from home, you first need to login to one of the jump hosts that are available for Internet access:

```console
ssh meling@ssh1.ux.uis.no
```

These additional jump hosts are available: `ssh2`, `ssh3`, and `ssh4`.
You may be required to use a 2FA code to log into these jump hosts.

If you are on campus however, you can skip the above step.
To log on to one of the Unix machines, I can run:

```console
ssh meling@gorina1.ux.uis.no
```

This may prompt for your password, and possibly a 2FA code.
Enter your account password and you will be logged into that machine.
However, if you are already logged in on one of the Unix machines, you should not need to type a password.

## Configuring SSH for Passwordless Login

In other environments than between the different Unix machines, you can avoid having to type the password each time by generating a public-private key-pair using the `ssh-keygen` command.
To see the man pages for `ssh-keygen` type:

```console
man ssh-keygen
```

and read the instructions.
Then try running this command to generate your key-pair; make sure that once asked to give a password, just press enter at the password prompt.
Once the key-pair have been generated, append the public-key file (ends with .pub) to a file named `authorized_keys`.

If you have multiple keys in the latter file, make sure not to overwrite those keys, and instead paste the new public-key at the end of your current file.
After having completed this process, try ssh to another machine and see whether you have to type the password again.

Note that the security of this passphrase-less method of authenticating towards a remote machine hinges on the protection of the private key file stored on your client machine.
Thus, it is actually recommended to create a key with a passphrase, and instead use the `ssh-agent` command at startup, along with `ssh-add` to add your key to this agent.
Then, the `ssh`, `scp`, and other ssh-based client commands can talk locally with the `ssh-agent`, and you as the user only need to type your password once.
Please consult the `ssh-agent` and `ssh-add` manual pages for additional details.

Tip: It is possible to create a `$HOME/.ssh/config` file, with content similar to this:

```config
Host uis
  HostName ssh4.ux.uis.no

Host gorina*
  HostName %h.ux.uis.no
  User username
  ProxyJump uis
```

Replace `username` with your own Unix username.
This should allow you to log directly into one of the `gorina` machines, e.g. by typing:

```console
ssh gorina1
```

You should also be able to log into the proxy jump machine, i.e., `ssh4.ux.uis.no` in the example above by typing:

```console
ssh uis
```

Note that, the proxy jump machines are not meant to be used for running lengthy jobs; it is okay to run simple short-lived commands on this machine.

To read more about setting up [SSH to remote hosts through a proxy or bastion with ProxyJump](https://www.redhat.com/sysadmin/ssh-proxy-bastion-proxyjump).

Another tip: If you are running from a laptop and wish to remain connected even if you close the laptop-lid, you can check out the [mosh command](http://mosh.mit.edu/).

### Task: Setting Up SSH Keys for Passwordless GitHub Authentication

In this task you will set up SSH keys on your local machine (or OpenStack VM) to enable passwordless authentication with GitHub.
This will allow you to clone repositories and push commits without entering your password each time.

#### Generating an SSH Key Pair

First, you need to generate an SSH key pair on your local machine or OpenStack VM.
Use `ssh-keygen` to generate a key pair as described in the previous sections.

```console
ssh-keygen -t ed25519 -C "your_email@example.com"
```

When prompted for a file location, press Enter to use the default location (`~/.ssh/id_ed25519`).

**Passphrase Options:**
You have two options when prompted for a passphrase:

1. **No passphrase (simpler)**: Press Enter to create a key without a passphrase.
   This simplifies the workflow as you won't need to enter anything when using the key.

2. **With passphrase (more secure)**: Enter a passphrase to protect your private key.
   For convenience, you can use `ssh-agent` to avoid typing the passphrase repeatedly:

   ```console
   # Start the ssh-agent (if not already running)
   eval "$(ssh-agent -s)"

   # Add your SSH key to the agent (you'll be prompted for the passphrase once)
   ssh-add ~/.ssh/id_ed25519
   ```

   After adding the key to the agent, you'll only need to enter the passphrase once per session.

For this course, either option works fine, but the no-passphrase option is simpler to get started.

#### Setting up SSH Authentication on GitHub

After generating your SSH key pair, you need to add the public key to your GitHub account.
Follow the instructions for [Connecting to GitHub with SSH](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh) including the step [Testing your SSH connection](https://docs.github.com/en/github/authenticating-to-github/testing-your-ssh-connection).
Note that these guides provide instructions for Mac, Windows and Linux.
You can select your OS via a tab near the top of each article.

## Cloning Your Assignments Repository and Submitting to QuickFeed

Carefully follow the instructions in the [lab submission guide](https://github.com/dat515-2025/info/blob/main/lab-submission.md).
The instructions should enable you to clone repositories and push changes to GitHub, fetch changes to the assignments, and submit your final solution.
This setup will work on your local machine, OpenStack VM, or any other environment where you've configured your SSH keys.
