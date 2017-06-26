# Tasks with different complexity

## Converter
Write a general-purpose unit-conversion program that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and k i log rams, and t he like. 

## String anagrams checker
Write a function that reports whet her two strings are anagrams of each other, that is, the y contain the same letters in a different order. 

## Frequency of each word in a text
Write a program wordfreq to report the frequency of each word in an input text file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines. 

## Github issue tool
Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required. 

## comic offline index
The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json pro duces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.

## tool poster
The JSON-based web service of the Open Movie Datab as e lets you s e arch https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line. 

## Crowl -> to local copy
Modify crawl to make local copies of the pages it finds, creating directories as necessary. Don’t make copies of pages that come from a different domain. For example, if the original page comes from golang.org, save all files from there, but exclude ones from vimeo.com. 

## CountingWriter
Write a function CountingWriter with the signature below that, given an io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer. 
  func CountingWriter(w io.Writer) (io.Writer, *int64) 

## LimitReader
The LimitReader function in the io package accepts an io.Reader r and a number of bytes n, and returns another Reader that reads from r but reports an end-of-file condition after n bytes. Implement it. 
     func LimitReader(r io.Reader, n int64) io.Reader

## IsPalindrome based on sort.Interface
The sort.Interface type can be adapted to other uses. Write a function IsPalindrome(s sort.Interface) bool that reports whet her the sequences is a palindrome, in other words, reversing the sequence would not change it. Assume that the elements at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i). 

## Chat server
Make the chat server disconnect idle clients, such as those that have sent no mess ages in the last five minutes. Hint: calling conn.Close() in another goroutine unblocks active Read calls such as the one done by input.Scan(). 

## Chat server2
Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it. Alter natively, add buffering to each client’s out going mess age channel so that most mess ages are not dropped; the broadcaster should use a non-blocking send to this channel. 

## Withdraw
Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program. The result should indicate whether the transaction succeeded or failed due to insufficient funds. The message sent to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to Withdraw. 

[SourceCode](https://github.com/adonovan/gopl.io/)

## JSON transformation
  input: {"id":"1234","forecast":{"temp":[14,15,26,25,16],"day":{"moon":[2,2,3,3,2],"wind":[23,24,17,17,25]}}}

  output: [{"id":"1234","forecast":{"temp":14,"day":{"moon":2,"wind":23}}},{"id":"1234","forecast":{"temp":15,"day":{"moon":2,"wind":24}}},{"id":"1234","forecast":{"temp":26,"day":{"moon":3,"wind":17}}},{"id":"1234","forecast":{"temp":25,"day":{"moon":3,"wind":17}}},{"id":"1234","forecast":{"temp":16,"day":{"moon":2,"wind":25}}}]

## API - time slots for distributed team
Create a secure API (HTTPS or HTTP/2) for helping distributed team find a time for meeting call. Endpoints:
- register new teamMember (admin rights)
- login (issue JWT) or JWT emitter
- time-zone (for some team member)
- schedule (setting available time slots, for some team member)
- meeting (return possible time slots for all team, all member should specify their time zone and schedule) 
