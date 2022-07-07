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
                options()[id].map(v => <span style={{ "grid-row": v / 3, "grid-column": v % 3 }}>{v}</span>)
            }
        </div>
    )
}