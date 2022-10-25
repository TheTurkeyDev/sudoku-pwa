/* @refresh reload */
import { render } from 'solid-js/web';

import './index.css';
import App from './App';
import { SudokuProvider } from './context/sudoku';
import { BoardType } from './board/board-type';

declare global {
    // eslint-disable-next-line @typescript-eslint/consistent-type-definitions
    interface Window {
        readonly generateBoard: (difficulty: number, callback: (error: string, board: string) => void) => void
    }
}

render(() => (
    <SudokuProvider>
        <App />
    </SudokuProvider>
), document.getElementById('root') as HTMLElement);
