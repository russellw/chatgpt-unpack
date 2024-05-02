ChatGPT allows you to download an archive of all the conversations you have had with it, so you can do things like search for past conversations with your choice of search tool.

The important file in the archive is `conversations.json`. It's not really usable directly, but `chatgpt-unpack` can unpack the data to a usable text file.

# Usage

Install Visual Studio, or the .Net SDK for the operating system you are using.

`cd` to the `chatgpt-unpack` directory.

Copy `conversations.json` to this directory.

Compile `chatgpt-unpack`. On Windows, this can be done from a Visual Studio command prompt, by just typing `msbuild`.

Run `chatgpt-unpack`.

On Windows, the supplied `unpack.bat` will both compile and run the program. Pull requests with similar scripts for other operating systems are welcome.

If all goes well, the program will generate `conversations.txt` containing a plain text version of the archive.
