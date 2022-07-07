import styles from './css/board.module.css';
import { Section } from './section';

export const Board = () => {
    return (
        <div class={styles.Board}>
            {
                Array.from(Array(9)).map((_, i) => <Section id={i} />)
            }
        </div>
    );
}