/*
You have a browser of one tab where you start on the homepage and you can visit another url, get back in the history number of steps or move forward in the history number of steps.

Implement the BrowserHistory class:

BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
void visit(string url) Visits url from the current page. It clears up all the forward history.
string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps. Return the current url after moving back in history at most steps.
string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x, you will forward only x steps. Return the current url after forwarding in history at most steps.


Example:

Input:
["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]
[["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]
Output:
[null,null,null,null,"facebook.com","google.com","facebook.com",null,"linkedin.com","google.com","leetcode.com"]

Explanation:
BrowserHistory browserHistory = new BrowserHistory("leetcode.com");
browserHistory.visit("google.com");       // You are in "leetcode.com". Visit "google.com"
browserHistory.visit("facebook.com");     // You are in "google.com". Visit "facebook.com"
browserHistory.visit("youtube.com");      // You are in "facebook.com". Visit "youtube.com"
browserHistory.back(1);                   // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
browserHistory.back(1);                   // You are in "facebook.com", move back to "google.com" return "google.com"
browserHistory.forward(1);                // You are in "google.com", move forward to "facebook.com" return "facebook.com"
browserHistory.visit("linkedin.com");     // You are in "facebook.com". Visit "linkedin.com"
browserHistory.forward(2);                // You are in "linkedin.com", you cannot move forward any steps.
browserHistory.back(2);                   // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
browserHistory.back(7);                   // You are in "google.com", you can move back only one step to "leetcode.com". return "leetcode.com"


Constraints:

1 <= homepage.length <= 20
1 <= url.length <= 20
1 <= steps <= 100
homepage and url consist of  '.' or lower case English letters.
At most 5000 calls will be made to visit, back, and forward.


Для решения данной задачи мы можем использовать структуру данных стек (stack) для хранения посещенных страниц и указатель текущей страницы.

Когда пользователь посещает новую страницу, мы добавляем ее в вершину стека и обнуляем все страницы в стеке после текущей страницы (т.е. удаляем все страницы, которые находятся ниже текущей страницы в стеке).

Когда пользователь нажимает кнопку "назад", мы просто перемещаем указатель текущей страницы на одну страницу вниз в стеке, если это возможно.

Когда пользователь нажимает кнопку "вперед", мы просто перемещаем указатель текущей страницы на одну страницу вверх в стеке, если это возможно.

Ниже представлен код на языке Golang, реализующий данную логику:

*/

package main

type BrowserHistory struct {
	stack []string // стек посещенных страниц
	curr  int      // указатель на текущую страницу
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (this *BrowserHistory) Visit(url string) {
	// добавляем новую страницу в стек и обнуляем все страницы после текущей
	this.stack = this.stack[:this.curr+1]
	this.stack = append(this.stack, url)
	this.curr = len(this.stack) - 1 // перемещаем указатель на новую страницу
}

func (this *BrowserHistory) Back(steps int) string {
	// перемещаем указатель на одну страницу назад в стеке, если это возможно
	this.curr = max(this.curr-steps, 0)
	return this.stack[this.curr]
}

func (this *BrowserHistory) Forward(steps int) string {
	// перемещаем указатель на одну страницу вперед в стеке, если это возможно
	this.curr = min(this.curr+steps, len(this.stack)-1)
	return this.stack[this.curr]
}

// функция для получения максимального значения из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// функция для получения минимального значения из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
