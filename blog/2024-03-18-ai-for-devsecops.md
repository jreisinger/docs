I've finally decided to try out a couple of AI related tools to see whether they are useful for me. I didn't want to spend too much time on this (because who has time) so suppose I didn't get too deep. I work in the Dev/Sec/Ops area, meaning I do small-scale programming (as opposed to full time application development), cybersecurity and IT operations. Since I use terminal a lot I had a look at three non-GUI tools. Here's what I've done and what are my conclusions so far.

## Ghatgpt-cli

First, I simply wanted a CLI interface to ChatGPT. One of the first Google results was [this project](https://github.com/kardolus/chatgpt-cli). You just need to give it your [API key](https://platform.openai.com/api-keys) either via environment variable (`export OPENAI_API_KEY=...`) or configuration file (enter `api_key: ...` to `~/.chatgpt-cli/config.yaml`) and you're ready to go. Now I don't have to open a browser window and log in into ChatGPT:

```
$ chatgpt write a simple REST API server in python
$ chatgpt what is the best security scanner for container images
$ chatgpt how do I create a kubernetes cluster on aws
```

I don't display the answers here to save paper but they are quite usable! Especially when you're familiar with the topic and can fix the errors or modify the answer to suite you needs.

## Fabric

I've been aware of Daniel Miessler's [project](https://github.com/danielmiessler/fabric) for a while. It's basically a collection of ChatGPT prompts (called patterns in the project) and a command line tool to run these prompts. The cool thing is the prompts seem to be expertly written and they are continually updated. Here are some of them. 

Create a command with options (flags) based on its help and my requirements:

```
$ tool=curl;echo -e "use $tool to silently download body to a file.\n\n$($tool -h 2>&1)" | fabric --pattern create_command
curl -s -o output.txt <url>

$ tool=wget;echo -e "use $tool to silently download body to a file.\n\n$($tool -h 2>&1)" | fabric --pattern create_command                                   
wget -q -O file URL
```

Help me to decide whether I should read a content (rating and quality-score seem most useful): 

```
$ wget -qO- https://jreisinger.blogspot.com/2022/09/preventing-git-leaks.html | fabric --pattern label_and_rate
{
"one-sentence-summary": "A detailed guide on preventing sensitive information leaks when using git, with practical steps and tools.",
"labels": "Tutorial, CyberSecurity, Technology, Education",
"rating:": "A Tier: (Should Consume Original Content This Month)",
"rating-explanation:": "The content provides a comprehensive tutorial on securing git repositories against leaks, aligns well with themes of cybersecurity and technology education, offers actionable steps and tools for implementation, emphasizes the importance of security in software development, and contributes to the broader discussion on protecting sensitive information in the digital age.",
"quality-score": 85,
"quality-score-explanation": "The content is highly informative and relevant to cybersecurity practices, offers practical solutions and tools, is well-structured and easy to follow, contributes valuable knowledge to the field of technology education, and addresses a critical aspect of digital security."
}
```

Summarize what a program does and explain its parts:

```
$ cat ~/github.com/jreisinger/tools/cmd/extracturl/main.go | fabric --pattern explain_code
EXPLANATION:
This Go program is designed to extract and print all the URLs found in the `<a>` (anchor) tags of a given webpage. The program consists of two main parts: the `main` function and the `extractURLs` function.

1. **Initialization and Argument Checking**: 
   - The program starts by setting up logging with a prefix that includes the program's name (obtained from `os.Args[0]`) and disabling additional flags for log output.
   - It then checks if a URL has been provided as a command-line argument. If not, it logs a fatal error message instructing the user to supply a URL.

2. **URL Extraction**:
   - The provided URL is passed to the `extractURLs` function.
   - This function attempts to make an HTTP GET request to the URL. If this request fails, it returns an error.
   - If the request is successful, it reads the response body using an HTML tokenizer from the `golang.org/x/net/html` package.
   - The tokenizer iterates through the HTML tokens in the response body. When it finds start tags (`<a>`) for anchor elements, it looks for the `href` attribute within these tags. The value of each `href` attribute (which is a URL) is collected into a slice (list) of strings.
   - Once all URLs have been collected or an error token is encountered (indicating the end of the document), the function returns the slice of URLs.

3. **Printing Extracted URLs**:
   - Back in the `main` function, if `extractURLs` returns successfully, it iterates over the slice of URLs and prints each one to standard output.

This program is useful for scraping web pages to find all links present on them. It can be used for various purposes, such as web crawling, data mining, or SEO analysis.
```

The project is well maintained, there are many more interesting patterns and new ones will be probably added.
 
## Beelzebub

I wrote about basic setup of this honeypot in a previous post. This time I wanted to see its AI part in action. The honeypot uses ChatGPT API to simulate a Linux terminal. After cloning [the repo](https://github.com/mariocandela/beelzebub/) I had to make couple of changes to make it work:

1. I added my ChatGPT API key to `configurations/services/ssh-2222.yaml`.
2. I changed the unsupported model:
```
diff --git a/plugins/openai-gpt.go b/plugins/openai-gpt.go
<...>
-               Model:            "text-davinci-003",
+               Model:            "gpt-3.5-turbo-instruct",
<...>
```
3. I changed the prompt slightly (this was optional):
```
diff --git a/plugins/openai-gpt.go b/plugins/openai-gpt.go
<...>
-       promptVirtualizeLinuxTerminal = "I want you to act as a Linux terminal. I will type commands and you will reply with what the terminal should show. I want you to only reply with the terminal output inside one unique code block, and nothing else. Do no write explanations. Do not type commands unless I instruct you to do so.\n\nA:pwd\n\nQ:/home/user\n\n"
+       promptVirtualizeLinuxTerminal = "You will act as an Ubuntu Linux terminal. The user will type commands, and you are to reply with what the terminal should show. Your responses must be contained within a single code block. Do not provide explanations or type commands unless explicitly instructed by the user. Remember previous commands and consider their effects on subsequent outputs.\n\nA:pwd\n\nQ:/home/user\n\n"
<...>
```

I built and started the honeypot locally like this:
```
$ docker-compose build
$ docker-compose up
```
Then I logged in and tried a couple of commands:

```
$ ssh root@localhost -p 2222
<...>
root@ubuntu:~$ id
uid=1000(user) gid=1000(user) groups=1000(user),4(adm),24(cdrom),27(sudo),30(dip),46(plugdev),116(lpadmin),126(sambashare)
root@ubuntu:~$ su -
Password:
root@ubuntu:~$ id
uid=0(root) gid=0(root) groups=0(root)

root@ubuntu:~$ cat /etc/shadow
root:*:17541:0:99999:7:::

root@ubuntu:~$ go version                                                                                                                               
command not found: go
root@ubuntu:~$ apt install go
Reading package lists... Done
root@ubuntu:~$ go version
go version go1.10.4 linux/amd64
```

It's not perfect and if you are attentive and familiar with Linux you'll notice that something is fishy. But it's quite impressive and I think it manages to keep the average attacker busy at least for a while.
