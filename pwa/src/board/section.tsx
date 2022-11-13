import { Cell } from './cell';
import styles from './css/section.module.css';

type SectionProps = {
    readonly id: number
}

export const Section = (props: SectionProps) => {
    const colStart = (props.id % 3) * 3;
    const rowStart = Math.floor(props.id / 3) * 3;
    return (
        <div class={styles.Section}>
            {Array.from(Array(9)).map((_, i) => <Cell id={((rowStart + Math.floor(i / 3)) * 9) + (colStart + (i % 3))} />)}
        </div>
    );
};