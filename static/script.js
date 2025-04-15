document.addEventListener("DOMContentLoaded", () => {
    const board = document.getElementById("board");
    const status = document.getElementById("status");
    const restartButton = document.getElementById("restart");

    function renderBoard(state) {
        board.innerHTML = "";
        state.board.forEach((cell, index) => {
            const div = document.createElement("div");
            div.classList.add("cell");
            div.dataset.index = index;

            // Если клетка не пустая, добавляем изображение
            if (cell === "X") {
                const img = document.createElement("img");
                img.src = "/static/cross.png";
                img.alt = "X";
                div.appendChild(img);
            } else if (cell === "O") {
                const img = document.createElement("img");
                img.src = "/static/zero.png";
                img.alt = "O";
                div.appendChild(img);
            }

            // Если это клетка из выигрышной линии, добавляем класс "winner"
            if (state.winner && state.winning_cells && state.winning_cells.includes(index)) {
                div.classList.add("winner");
            }

            div.addEventListener("click", handleClick);
            board.appendChild(div);
        });

        status.textContent = state.winner ? `Победитель: ${state.winner}` : `Ходит: ${state.turn}`;
    }

    function showLoading() {
        status.textContent = "Загрузка...";
    }

    function fetchState() {
        showLoading();
        fetch("/game/", {
            method: "GET"
        })
            .then(res => {
                if (!res.ok) throw new Error(`HTTP error: ${res.status}`);
                return res.json();
            })
            .then(renderBoard)
            .catch(err => {
                console.error("Error fetching state:", err);
                status.textContent = "Ошибка загрузки";
            });
    }

    function handleClick(event) {
        fetch("/game/", { method: "GET" })
            .then(res => {
                if (!res.ok) throw new Error(`HTTP error: ${res.status}`);
                return res.json();
            })
            .then(state => {
                if (state.winner) return;
                const index = event.target.dataset.index;
                showLoading();
                fetch("/game/", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ index: parseInt(index) })
                })
                    .then(res => {
                        if (!res.ok) throw new Error(`HTTP error: ${res.status}`);
                        return res.json();
                    })
                    .then(renderBoard)
                    .catch(err => {
                        console.error("Error making move:", err);
                        status.textContent = "Ошибка хода";
                    });
            })
            .catch(err => console.error("Error checking game state:", err));
    }

    restartButton.addEventListener("click", () => {
        showLoading();
        fetch("/game/", {
            method: "DELETE"
        })
            .then(res => {
                if (!res.ok) throw new Error(`HTTP error: ${res.status}`);
                return res.json();
            })
            .then(renderBoard)
            .catch(err => {
                console.error("Error resetting game:", err);
                status.textContent = "Ошибка сброса";
            });
    });

    fetchState();
});


