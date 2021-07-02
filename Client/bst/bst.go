/*
Package bst is the binarytree data structure for the API data to be added to it so the client can
retrieve data from the BST instead of constantly querying the API
*/
package bst

import (
	"PGL/Client/models"
	"sync"
)

//Node structure containing the user information
type Node struct {
	User  models.User
	Cats  []models.Category
	left  *Node
	right *Node
}

//BST
type BST struct {
	root  *Node
	lock  sync.RWMutex
	Count int
}

//add with recursive
func (b *BST) add(n **Node, user models.User, cats []models.Category) {
	if *n == nil {
		newNode := &Node{
			User:  user,
			Cats:  cats,
			left:  nil,
			right: nil,
		}
		*n = newNode
	}
	if user.Username < (*n).User.Username {
		b.add(&((*n).left), user, cats)
	} else {
		b.add(&((*n).right), user, cats)
	}
}

//add wrapper
func (b *BST) Add(user models.User, cats []models.Category) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.add(&b.root, user, cats)
	b.Count++
}

//delete with recursive
func (b *BST) delete(n **Node, username string) *Node {
	if *n == nil {
		return nil
	}
	if username < (*n).User.Username {
		(*n).left = b.delete(&((*n).left), username)
		return *n
	}
	if username > (*n).User.Username {
		(*n).right = b.delete(&((*n).right), username)
		return *n
	}

	//case 1: node has no sub-trees
	if (*n).left == nil && (*n).right == nil {
		*n = nil
		return nil
	}
	//case 2: node has 1 sub-tree
	if (*n).left == nil { // case 2 : node has only 1 sub-tree
		*n = (*n).right
		return *n
	}
	if (*n).right == nil { // case 2 : node has  only 1 sub-tree
		*n = (*n).left
		return *n
	}
	//case 3: node 2 has 2 sub-trees
	current := (*n).left
	for {

		if current != nil && current.right != nil {
			current = current.right
		} else {
			break
		}
	}
	(*n).User.Username = current.User.Username
	(*n).left = b.delete(&(*n).left, current.User.Username)
	return *n
}

//delete wrapper
func (b *BST) Delete(username string) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.delete(&b.root, username)
	b.Count--
}

//get with recursive
func (b *BST) get(n *Node, username string) *Node {
	if n == nil {
		return nil
	} else if username == n.User.Username {
		return n
	} else if username < n.User.Username {
		return b.get(n.left, username)
	} else {
		return b.get(n.right, username)
	}
}

//get wrapper
func (b *BST) Get(username string) Node {
	b.lock.RLock()
	defer b.lock.RUnlock()
	info := *b.get(b.root, username)
	return info
}
