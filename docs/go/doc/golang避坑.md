// bad
names := []string{"andy", "ben", "cindy"}
namesRef := make([]*string, len(names))
for idx, name := range names {
   namesRef[idx] = &name
}
for _, nameRef := range namesRef {
   fmt.Println(*nameRef)
}
// output: cindy, cindy, cindy
 
// good
names := []string{"andy", "ben", "cindy"}
namesRef := make([]*string, len(names))
for idx := range names { // only allow to use loop index
   namesRef[idx] = &names[idx]
}
for _, nameRef := range namesRef {
   fmt.Println(*nameRef)
}