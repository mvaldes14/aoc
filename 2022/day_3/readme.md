# Notes

In this exercise the trick is to find the numerical value of each letter in the alphabet so you can then just do a sum of the ones that are in either sack.

In go, each letter is a UINT8 which has a numerical value on each string it parses. For example the letter A is 97.

You can also do math operations on strings unlike some other languages so if you do.

```golang
fmt.Println('a' - 'a')

```
It prints out 0 which matches the first letter of the alphabet.
So by that logic if you do `'b' - 'a'` you get 1, meaning its the 2nd letter in the alphabet.

If you iterate over the entire thing you can find that the first lowercase letters can be obtained by doing `'letter' - 'a'` to get its actual value.

Same logic applies to uppercase letters but since they are next in the alphabet you need to sum 26, so it represents the next set. The actual value of the string changes depending if its upper or lowercase.

```golang
fmt.Println('a') //97
fmt.Println('A') //65
```

