import { useSudoku } from '../context/sudoku';
import styles from './css/options.module.css';

type OptionsProps = {
    readonly id: number
}

export const Options = (props: OptionsProps) => {
    const { options } = useSudoku();

    return (
        <div class={styles.Options}>
            {
                [1, 2, 3, 4, 5, 6, 7, 8, 9].map(v => (
                    <span>{options()[props.id].includes(v) ? v : ''}</span>
                ))
            }
        </div>
    );
};