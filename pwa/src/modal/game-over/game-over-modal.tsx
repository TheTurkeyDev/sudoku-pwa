import { createEffect, createSignal } from 'solid-js';
import { useSudoku } from '../../context/sudoku';
import { timeBetweenDates } from '../../header/timer';
import { Modal } from '../modal';
import styles from './game-over-modal.module.css';

export const GameOverModal = () => {
    const { gameState, setGameState, setDifficulty, startTime, endTime } = useSudoku();
    const [time, setTime] = createSignal<string>('');

    createEffect(() => {
        const data = timeBetweenDates(startTime(), endTime());
        setTime(`${data.minutes}:${data.seconds}`);
    });


    return (
        <Modal show={gameState() === 2}>
            <div class={styles.GameOverWrapper}>
                <h1>Game Over!</h1>
                <span>Completion Time: {time()}</span>
                <div class={styles.ButtonRow}>
                    <button onClick={() => { setGameState(0); setDifficulty(-1); }}>Main Menu</button>
                </div>
            </div>
        </Modal>
    );
};