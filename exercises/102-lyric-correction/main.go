package main

import (
	"fmt"
	"strings"
)

func main() {
	// correct the lyrics
	// before [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
	// after [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]

	// before
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// prepend 'yesterday'
	lyric = append([]string{"yesterday"}, lyric...)

	// rearrange by swapping second and third phrases
	second := append([]string(nil), lyric[13:]...)
	third := append([]string(nil), lyric[8:13]...)

	lyric = append(lyric[:8], second...)
	lyric = append(lyric, third...)
	fmt.Println(lyric)

	fmt.Printf("'lyric' slice Length: %d Capacity: %d\n", len(lyric), cap(lyric))

	// alternative solution
	lyric2 := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// prepend 'yesterday'
	lyric2 = append([]string{"yesterday"}, lyric2...)

	// extract 'oh I believe in yesterday' and append it to the end
	lyric2 = append(lyric2, lyric2[8:13]...)
	// remove the middle 'oh I believe in yesterday' by overwriting
	lyric2 = append(lyric2[:8], lyric2[13:]...)

	fmt.Println(lyric2)

	fmt.Printf("'lyric2' slice Length: %d Capacity: %d\n", len(lyric2), cap(lyric2))
}
