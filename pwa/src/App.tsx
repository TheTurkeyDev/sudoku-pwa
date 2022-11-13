import { createEffect, createResource, createSignal } from 'solid-js';
import styles from './App.module.css';
import { BoardType } from './board/board-type';
import { useSudoku } from './context/sudoku';
import { CenteredLoadingSpinner, LoadingSpinner } from './loading/loading-spinner';
import { NavBar } from './navbar/navbar';
import { SudokuPuzzle } from './sudoku-puzzle';

const generateBoardPromise = (difficulty: number): Promise<string> => (
    new Promise<string>((resolve, reject) => {
        if (difficulty === -1) {
            resolve('');
            return;
        }

        //It hurts
        setTimeout(() => {
            window.generateBoard(difficulty, (err, message) => {
                if (err) {
                    reject(err);
                    return;
                }

                resolve(message);
            });
        }, 100);
    })
);

const App = () => {
    const { loadBoard, setDifficulty, difficulty, gameState } = useSudoku();
    const [boardJson] = createResource<string, number>(difficulty, generateBoardPromise);

    createEffect(() => {
        if (difficulty() === -1)
            return;
        const json = boardJson();
        if (!json)
            return;
        const board = JSON.parse(json) as BoardType;
        loadBoard(board);
    });

    const playDailyLevel = () => {
        //console.log(window.generateBoard(difficulty))
    };

    return (
        <div class={styles.App}>
            <NavBar />
            {
                boardJson.loading ? <CenteredLoadingSpinner /> : (gameState() > 0 ? <SudokuPuzzle /> : (
                    <div class={styles.MainContent}>
                        <h1>Daily Sudoku</h1>
                        <hr style={{ width: '100%' }} />
                        <button onClick={playDailyLevel}>Daily Puzzle</button>
                        <button onClick={() => setDifficulty(0)}>Beginner</button>
                        <button onClick={() => setDifficulty(1)}>Easy</button>
                        <button onClick={() => setDifficulty(2)}>Medium</button>
                        <button onClick={() => setDifficulty(3)}>Hard</button>
                    </div>
                ))
            }
        </div >
    );
};

export default App;
