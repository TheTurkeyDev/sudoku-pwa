/* @refresh reload */
import { render } from 'solid-js/web';

import './index.css';
import App from './App';
import { SudokuProvider } from './context/sudoku';

declare global {
    interface Window {
        generateBoard: (difficulty: number, callback: (error: string, board: any) => void) => void
    }
}

render(() => (
    <SudokuProvider>
        <App />
    </SudokuProvider>
), document.getElementById('root') as HTMLElement);
