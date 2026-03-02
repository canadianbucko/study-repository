import sys
from PyQt5.QtWidgets import QApplication, QMainWindow, QPushButton, QLabel
from PyQt5.QtCore import Qt


class TicTacToeGame(QMainWindow):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("Крестики нолики pyqt5")
        self.setGeometry(100, 100, 300, 360)

        self.buttons = []
        self.current_player = "X"
        self.move_count = 0
        self.game_over = False

        self.create_buttons()

        self.status_label = QLabel(self)
        self.status_label.setGeometry(0, 300, 300, 20)
        self.status_label.setAlignment(Qt.AlignCenter)
        self.status_label.setText("Ход: X")

        self.new_game_btn = QPushButton("Новая игра", self)
        self.new_game_btn.setGeometry(0, 320, 300, 40)
        self.new_game_btn.clicked.connect(self.reset_game)

    def create_buttons(self):
        for row in range(3):
            row_buttons = []
            for col in range(3):
                button = QPushButton(self)
                button.setText("")
                button.setGeometry(col * 100, row * 100, 100, 100)
                button.clicked.connect(self.on_click)
                row_buttons.append(button)
            self.buttons.append(row_buttons)

    def on_click(self):
        if self.game_over:
            return

        button = self.sender()
        if button.text() == "":
            button.setText(self.current_player)
            self.move_count += 1

            # чек на победу/ничью
            if self.check_winner():
                return

            # меняем игрока если не кончилась
            self.switch_player()
            self.status_label.setText("Ход: " + self.current_player)

    def check_winner(self):
        for r in range(3):
            if self.buttons[r][0].text() == self.buttons[r][1].text() == self.buttons[r][2].text() != "":
                self.end_game("Победил игрок: " + self.buttons[r][0].text())
                return True


        for c in range(3):
            if self.buttons[0][c].text() == self.buttons[1][c].text() == self.buttons[2][c].text() != "":
                self.end_game("Победил игрок: " + self.buttons[0][c].text())
                return True

        if self.buttons[0][0].text() == self.buttons[1][1].text() == self.buttons[2][2].text() != "":
            self.end_game("Победил игрок: " + self.buttons[0][0].text())
            return True

        if self.buttons[0][2].text() == self.buttons[1][1].text() == self.buttons[2][0].text() != "":
            self.end_game("Победил игрок: " + self.buttons[0][2].text())
            return True

        if self.move_count == 9:
            self.end_game("Ничья")
            return True

        return False

    def end_game(self, message):
        self.game_over = True
        self.status_label.setText(message)
        for row in self.buttons:
            for button in row:
                button.setEnabled(False)

    def switch_player(self):
        if self.current_player == "X":
            self.current_player = "O"
        else:
            self.current_player = "X"

    def reset_game(self):
        self.game_over = False
        self.current_player = "X"
        self.move_count = 0
        self.status_label.setText("Ход: X")

        for row in self.buttons:
            for button in row:
                button.setText("")
                button.setEnabled(True)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = TicTacToeGame()
    window.show()
    sys.exit(app.exec_())
