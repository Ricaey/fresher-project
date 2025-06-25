// filepath: comment-section-web/comment-section-web/src/scripts/main.js
document.addEventListener('DOMContentLoaded', function() {
    const commentForm = document.getElementById('comment-form');
    const commentInput = document.getElementById('comment-input');
    const commentList = document.getElementById('comment-list');

    commentForm.addEventListener('submit', function(event) {
        event.preventDefault();
        
        const commentText = commentInput.value.trim();
        if (commentText) {
            const commentItem = document.createElement('div');
            commentItem.classList.add('comment-item');
            commentItem.textContent = commentText;
            commentList.appendChild(commentItem);
            commentInput.value = '';
        }
    });
});