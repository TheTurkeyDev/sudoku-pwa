import styles from './sudoku-puzzle.module.css';
import { useSudoku } from './context/sudoku';
import { Board } from './board/board';
import { Keyboard } from './keyboard/keyboard';

export const SudokuPuzzle = () => {
  const { setBoardValue, setSelectedCell, selectedCell, onInput } = useSudoku();

  const keyDownHandler = (event: KeyboardEvent) => {
    if (event.code.startsWith('Digit')) {
      const digit = parseInt(event.code.replace('Digit', ''));
      if (isNaN(digit) || digit < 1 || digit > 9)
        return;
      onInput(digit);
    }

    switch (event.code) {
      case 'Backspace':
        setBoardValue(selectedCell(), 0);
        break;
      case 'ArrowUp':
        setSelectedCell(Math.max(selectedCell() % 9, selectedCell() - 9));
        break;
      case 'ArrowDown':
        setSelectedCell(Math.min(72 + (selectedCell() % 9), selectedCell() + 9));
        break;
      case 'ArrowLeft':
        setSelectedCell(Math.max(Math.floor(selectedCell() / 9) * 9, selectedCell() - 1));
        break;
      case 'ArrowRight':
        setSelectedCell(Math.min((Math.ceil((selectedCell() + 1) / 9) * 9) - 1, selectedCell() + 1));
        break;
    }
  }

  return (
    <div class={styles.Content} tabIndex={0} onKeyDown={keyDownHandler}>
      <Board />
      <Keyboard />
    </div>
  );
}