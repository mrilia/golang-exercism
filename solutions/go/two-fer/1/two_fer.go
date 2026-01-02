
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
    msgPrefix := "One for "
 	msgPostfix := ", one for me."
    
	return msgPrefix + name + msgPostfix
}
