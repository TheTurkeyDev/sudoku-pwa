import { useSudoku } from '../context/sudoku';
import styles from './css/cell.module.css';
import { Options } from './options';

type CellProps = {
    id: number
}
export const Cell = ({ id }: CellProps) => {
    const { board, setSelectedCell, selectedCell } = useSudoku();

    return (
        <div class={`${styles.Cell} ${selectedCell() === id ? styles.CellSelected : ''}`} onClick={() => setSelectedCell(id)}>
            {
                board()[id] == 0 ? <Options id={id} /> : <h2 class={styles.Number}>{board()[id]}</h2>
            }
        </div>
    );
}