import styles from './App.module.css';
import { Board } from './board/board';
import { NavBar } from './navbar/navbar';

const App = () => {
  return (
    <div class={styles.App}>
      <NavBar />
      <Board />
    </div>
  );
};

export default App;
