import { createEffect, createSignal } from 'solid-js';
import { useSudoku } from '../context/sudoku';
import styles from './css/cell.module.css';
import { Options } from './options';

type CellProps = {
    id: number
}
export const Cell = ({ id }: CellProps) => {
    const { board, setSelectedCell, selectedCell, isLockedValue } = useSudoku();
    const [highlight, setHighlighted] = createSignal<boolean>();

    const locked = isLockedValue(id)

    createEffect(() => {
        const selectedCellVal = board()[selectedCell()];
        setHighlighted(selectedCellVal !== 0 && selectedCellVal === board()[id]);
    })

    return (
        <div
            class={`${styles.Cell} ${selectedCell() === id ? styles.CellSelected : ''} ${highlight() ? styles.CellHighlight : ''}`}
            onClick={() => setSelectedCell(id)}>
            {
                board()[id] == 0 ? <Options id={id} /> : (
                    <h2 class={`${styles.Number} ${locked ? '' : styles.Unlocked}`}>
                        {board()[id]}
                    </h2>
                )
            }
        </div>
    );
}