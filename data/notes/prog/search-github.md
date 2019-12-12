Searching Your Github
---------------------

I have some code and documetation hosted on GitHub and I sometimes need to search through all these data.

1. Use [gitmeta](https://github.com/jreisinger/gitmeta) to clone/update the desired repos.

2. Search through the data

        # Look for files
        find -type f -iname "*back*"
    
        # Search inside files
        grep --color -R backup
