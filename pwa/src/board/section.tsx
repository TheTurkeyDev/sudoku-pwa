import { Cell } from './cell';
import styles from './css/section.module.css';

export const Section = () => {
    return (
        <div class={styles.Section}>
            {Array.from(Array(9)).map(() => <Cell />)}
        </div>
    );
}