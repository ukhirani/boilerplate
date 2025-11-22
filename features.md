
the set of features that are to be implemented in the first iteration

documentation


# basic commands

for anyone to get started and to know the in's and out's of the utility

bp - small intro of what the command actually does
(similar to bp --help or bp -H)
bp --help or bp -H (obviously will display all the things that the command does)

- bp list or bp ls (show all the templates that i can reproduce right now)

# core commands

- bp "template-name" "file-name"
the core of this whole utility, to take a name u remember for a template and make a file/directory 
that will have the things you planned for

- bp --preview -P "template-name" 
show the contents within the template that you want to preview

# how to add your own templates

- bp add "file-name"
it will add that file to the template repository with the template name as the file name 
- bp add "file-name" --tag "template-name"
same as the command above, just the only thing is that it will consider the tag as the template-name rather than the file name
