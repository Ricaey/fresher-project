const API_HOST = 'http://localhost:8080'; // 替换为你的实际host

document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.getElementById('login-form');
    const usernameInput = document.getElementById('username-input');
    const welcomeDiv = document.getElementById('welcome');
    const currentUsernameSpan = document.getElementById('current-username');
    const logoutBtn = document.getElementById('logout-btn');
    const commentForm = document.getElementById('comment-form');
    const commentInput = document.getElementById('comment-input');
    const commentList = document.getElementById('comment-list');

    let currentUser = localStorage.getItem('username') || '';

    function updateUI() {
        if (currentUser) {
            loginForm.style.display = 'none';
            welcomeDiv.style.display = '';
            currentUsernameSpan.textContent = currentUser;
            commentForm.style.display = '';
        } else {
            loginForm.style.display = '';
            welcomeDiv.style.display = 'none';
            commentForm.style.display = 'none';
        }
        fetchComments();
    }

    // 获取评论
    function fetchComments() {
        fetch(`${API_HOST}/comment/get?page=1&size=-1`)
            .then(res => res.json())
            .then(resp => {
                if (resp.code === 0) {
                    renderComments(resp.data.comments || []);
                } else {
                    commentList.innerHTML = `<div style="color:red;">获取评论失败：${resp.msg}</div>`;
                }
            })
            .catch(() => {
                commentList.innerHTML = `<div style="color:red;">无法连接服务器</div>`;
            });
    }

    // 渲染评论
    function renderComments(comments) {
        commentList.innerHTML = '';
        comments.forEach(comment => {
            const commentItem = document.createElement('div');
            commentItem.classList.add('comment-item');
            commentItem.innerHTML = `
                <span class="comment-author">${comment.name}：</span>
                <span class="comment-text">${comment.content}</span>
            `;
            // 只有本人可以删除
            if (currentUser && comment.name === currentUser) {
                const delBtn = document.createElement('button');
                delBtn.textContent = '删除';
                delBtn.className = 'delete-btn';
                delBtn.onclick = function() {
                    deleteComment(comment.id);
                };
                commentItem.appendChild(delBtn);
            }
            commentList.appendChild(commentItem);
        });
    }

    // 添加评论
    commentForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const commentText = commentInput.value.trim();
        if (commentText && currentUser) {
            fetch(`${API_HOST}/comment/add`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: currentUser,
                    content: commentText
                })
            })
            .then(res => res.json())
            .then(resp => {
                if (resp.code === 0) {
                    commentInput.value = '';
                    fetchComments();
                } else {
                    alert('评论失败：' + resp.msg);
                }
            })
            .catch(() => alert('无法连接服务器'));
        }
    });

    // 删除评论
    function deleteComment(id) {
        fetch(`${API_HOST}/comment/delete?id=${id}`, {
            method: 'POST'
        })
        .then(res => res.json())
        .then(resp => {
            if (resp.code === 0) {
                fetchComments();
            } else {
                alert('删除失败：' + resp.msg);
            }
        })
        .catch(() => alert('无法连接服务器'));
    }

    // 登录
    loginForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const username = usernameInput.value.trim();
        if (username) {
            currentUser = username;
            localStorage.setItem('username', username);
            updateUI();
        }
    });

    // 登出
    logoutBtn.addEventListener('click', function() {
        localStorage.removeItem('username');
        currentUser = '';
        updateUI();
    });

    updateUI();
});
