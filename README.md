# Angelou - Generate the Rhyme Scheme for any Poem in a flash

Simply pass in the name of a file with the poem and watch as the scheme is saved for you in a file named ./rhyme_scheme.txt

## Algorithm

- Reads a given poem into memory
- Separate the poem into lines
- Distrubutes the lines to a workerpool
- Each Worker gets the last syllable for each line
- Match up the last syllables and place them in a string
- Save the rhyme scheme to a .txt file

## Issues

- Angelou only works for poems with up to 52 different ending syllables (A-Za-z)

### Thank you to https://golangbot.com/buffered-channels-worker-pools/ for teaching me about worker pools