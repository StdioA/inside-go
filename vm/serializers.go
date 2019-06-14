package vm

import "github.com/stdioa/inside-go/db"

func SerializePost(post *db.Post) Post {
	postVM := Post{
		post.ID,
		post.Content,
		post.CreatedAt,
		SerializeComments(post.Comments),
	}
	return postVM
}

func SerializePosts(posts []db.Post) []Post {
	postVMs := make([]Post, 0, len(posts))
	for _, p := range posts {
		postVMs = append(postVMs, SerializePost(&p))
	}
	return postVMs
}

func SerializeComment(comment *db.Comment) Comment {
	return Comment{
		Content: comment.Content,
		Author:  comment.Author,
	}
}

func SerializeComments(comments []db.Comment) []Comment {
	commentVMs := make([]Comment, 0, len(comments))
	for _, c := range comments {
		commentVMs = append(commentVMs, SerializeComment(&c))
	}
	return commentVMs
}
