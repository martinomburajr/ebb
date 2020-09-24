package tree

var (
	BTree0Exp = BinaryTree{{key: 0, value: '0'}}
	BTree0    = BinaryTree{{key: 0, value: '0'}}
	BTree0xX  = BinaryTree{{key: 0, value: '('}, {key: 1, value: '0'}, {key: 2, value: ')'}}

	BTree1Exp = BinaryTree{{key: 0, value: 'x'}, {key: 0, value: '+'}, {key: 0, value: '1'}}
	BTree1    = BinaryTree{{key: 0, value: '('}, {key: 1, value: 'x'}, {key: 2, value: '+'}, {key: 3, value: '1'}, {key: 4, value: ')'}}

	BTree2Exp = BinaryTree{{key: 0, value: 'x'}, {key: 0, value: '*'}, {key: 0, value: 'x'}, {key: 0, value: '+'}, {key: 0, value: '1'}}
	BTree2    = BinaryTree{{key: 0, value: '('}, {key: 1, value: '('}, {key: 2, value: 'x'}, {key: 3, value: '*'}, {key: 4, value: 'x'}, {key: 5, value: ')'}, {key: 6, value: '+'}, {key: 7, value: '1'}, {key: 8, value: ')'}}
	// BTree1 in BTree2 at Position (index) 0
	BTree2_BTree1_NTP0 = BinaryTree{{key: 0, value: '('}, {key: 1, value: '('}, {key: 2, value: 'x'}, {key: 3, value: '+'}, {key: 4, value: '1'}, {key: 5, value: ')'}, {key: 6, value: '+'}, {key: 7, value: '1'}, {key: 8, value: ')'}}

	BTree3Exp           = BinaryTree{{key: 0, value: 'x'}, {key: 0, value: '*'}, {key: 0, value: 'x'}, {key: 0, value: '+'}, {key: 0, value: '1'}, {key: 0, value: '/'}, {key: 0, value: '3'}}
	BTree3              = BinaryTree{{key: 0, value: '('}, {key: 1, value: '('}, {key: 2, value: '('}, {key: 3, value: 'x'}, {key: 4, value: '*'}, {key: 5, value: 'x'}, {key: 6, value: ')'}, {key: 7, value: '+'}, {key: 8, value: '1'}, {key: 9, value: ')'}, {key: 10, value: '/'}, {key: 11, value: '3'}, {key: 12, value: ')'}}
	BTree3__BTree1_NTP0 = BinaryTree{{key: 0, value: '('}, {key: 1, value: '('}, {key: 2, value: '('}, {key: 3, value: 'x'}, {key: 4, value: '+'}, {key: 5, value: '1'}, {key: 6, value: ')'}, {key: 7, value: '+'}, {key: 8, value: '1'}, {key: 9, value: ')'}, {key: 10, value: '/'}, {key: 11, value: '3'}, {key: 12, value: ')'}}
	BTree3__BTree1_NTP1 = BinaryTree{{key: 0, value: '('}, {key: 1, value: '('}, {key: 2, value: 'x'}, {key: 3, value: '+'}, {key: 4, value: '1'}, {key: 5, value: ')'}, {key: 6, value: '/'}, {key: 7, value: '3'}, {key: 8, value: ')'}}

	// ((((x*x)+1)/3)-6)
	BTree4 = BinaryTree{
		{key: 0, value: '('},
		{key: 1, value: '('},
		{key: 2, value: '('},
		{key: 3, value: '('},
		{key: 4, value: 'x'},
		{key: 5, value: '*'},
		{key: 6, value: 'x'},
		{key: 7, value: ')'},
		{key: 8, value: '+'},
		{key: 9, value: '1'},
		{key: 10, value: ')'},
		{key: 11, value: '/'},
		{key: 12, value: '3'},
		{key: 13, value: ')'},
		{key: 14, value: '-'},
		{key: 15, value: '6'},
		{key: 16, value: ')'},
	}

	// Terminals
	BTreeTerminals    = []rune{'x', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	BTreeNonTerminals = []rune{'*', '+', '-', '/'}
)
