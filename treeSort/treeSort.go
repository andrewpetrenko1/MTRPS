package treeSort

type node struct {
	value string
	left  *node
	right *node
}

type bst struct {
	root *node
}

func (t *bst) insert(v string) {
	if t.root == nil {
		t.root = &node{v, nil, nil}
		return
	}
	current := t.root
	for {
		if v < current.value {
			if current.left == nil {
				current.left = &node{v, nil, nil}
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &node{v, nil, nil}
				return
			}
			current = current.right
		}
	}
}

func (t *bst) insertRev(v string) {
	if t.root == nil {
		t.root = &node{v, nil, nil}
		return
	}
	current := t.root
	for {
		if v > current.value {
			if current.left == nil {
				current.left = &node{v, nil, nil}
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &node{v, nil, nil}
				return
			}
			current = current.right
		}
	}
}

func (t *bst) inorder(visit func(string)) {
	var traverse func(*node)
	traverse = func(current *node) {
		if current == nil {
			return
		}
		traverse(current.left)
		visit(current.value)
		traverse(current.right)
	}
	traverse(t.root)
}

func (t *bst) slice() []string {
	sliced := []string{}
	t.inorder(func(v string) {
		sliced = append(sliced, v)
	})
	return sliced
}

func Treesort(values []string, reverse bool) []string {
	tree := bst{}
	for _, v := range values {
    if reverse {
      tree.insertRev(v)
    } else {
		  tree.insert(v)
    }
	}
	return tree.slice()
}