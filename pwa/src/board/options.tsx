import { useSudoku } from '../context/sudoku';
import styles from './css/options.module.css';

type OptionsProps = {
    id: number
}

export const Options = ({ id }: OptionsProps) => {
    const { options } = useSudoku();

    return (
        <div class={styles.Options}>
            {
                [1, 2, 3, 4, 5, 6, 7, 8, 9].map(v => (
                    <span>{options()[id].includes(v) ? v : ''}</span>
                ))
            }
        </div>
    )
}