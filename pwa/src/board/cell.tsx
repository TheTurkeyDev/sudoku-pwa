import styles from './css/cell.module.css';
import { Options } from './options';

export const Cell = () => {
    const showOptions = false;

    return (
        <div class={styles.Cell}>
            {
                showOptions ? <Options /> : <h2 class={styles.Number}>1</h2>
            }
        </div>
    );
}