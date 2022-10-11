import styles from './App.module.css';
import { Board } from './board/board';
import { useSudoku } from './context/sudoku';
import { Keyboard } from './keyboard/keyboard';
import { NavBar } from './navbar/navbar';

const App = () => {
  const { setBoardValue, selectedCell, onInput } = useSudoku();

  const keyDownHandler = (event: KeyboardEvent) => {
    if (event.code.startsWith('Digit')) {
      const digit = parseInt(event.code.replace('Digit', ''));
      if (isNaN(digit) || digit < 1 || digit > 9)
        return;
      onInput(digit);
    }
    else if (event.code === 'Backspace') {
      setBoardValue(selectedCell(), 0);
    }
  }

  console.log(window.generateBoard(1))

  return (
    <div class={styles.App} tabIndex={0} onKeyPress={keyDownHandler}>
      <NavBar />
      <Board />
      <Keyboard />
    </div>
  );
};

export default App;
