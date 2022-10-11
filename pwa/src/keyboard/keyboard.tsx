import { useSudoku } from '../context/sudoku';
import styles from './keyboard.module.css';

export const Keyboard = () => {
    const { setBoardValue, onInput, selectedCell, setEditingOptions, editingOptions } = useSudoku();

    const onKeyClick = (key: number) => {
        if (isNaN(key) || key < 1 || key > 9)
            return;
        onInput(key);
    }

    const deleteValue = () => {
        setBoardValue(selectedCell(), 0);
    }

    return (
        <div class={styles.KeyboardWrapper}>
            {[1, 2, 3, 4, 5, 6, 7, 8, 9].map(k => (
                <button class={styles.Key} onClick={() => onKeyClick(k)}>{k}</button>
            ))}
            <button class={styles.Key} onClick={deleteValue}>←</button>
            <button class={`${styles.Key} ${editingOptions() ? styles.KeySelected : ''}`} onClick={() => setEditingOptions(!editingOptions())}>✎</button>
        </div>
    );
}