import { createEffect, createSignal } from 'solid-js';
import { useSudoku } from '../context/sudoku';
import styles from './css/cell.module.css';
import { Options } from './options';

type CellProps = {
    readonly id: number
}

const doValidCheck = false;

export const Cell = (props: CellProps) => {
    const { board, setSelectedCell, selectedCell, isLockedValue, solution } = useSudoku();
    const [highlight, setHighlighted] = createSignal<boolean>();
    const [invalid, setInvalid] = createSignal<boolean>();

    const locked = isLockedValue(props.id);

    createEffect(() => {
        const selectedCellVal = board()[selectedCell()];
        setHighlighted(selectedCellVal !== 0 && selectedCellVal === board()[props.id]);
    });

    const checkVert = (id: number, val: number) => !!board().find((v, i) => i !== id && i % 9 === id % 9 && v === val);
    const checkHoriz = (id: number, val: number) => !!board().find((v, i) => i !== id && Math.floor(i / 9) === Math.floor(id / 9) && v === val);
    const checkBox = (id: number, val: number) => !!board().find((v, i) => i !== id && Math.floor(i / 27) === Math.floor(id / 27) && Math.floor((i % 9) / 3) === Math.floor((id % 9) / 3) && v === val);

    createEffect(() => {
        const cellValue = board()[props.id];
        if (doValidCheck)
            setInvalid(!!cellValue && solution()[props.id] !== cellValue);
        else
            setInvalid(checkVert(props.id, cellValue) || checkHoriz(props.id, cellValue) || checkBox(props.id, cellValue));
    });

    return (
        <div
            class={`${styles.Cell} ${selectedCell() === props.id ? styles.CellSelected : ''} ${highlight() ? styles.CellHighlight : ''}`}
            onClick={() => setSelectedCell(props.id)}>
            {
                board()[props.id] === 0 ? <Options id={props.id} /> : (
                    <h2 class={`${styles.Number} ${locked ? '' : styles.Unlocked} ${!locked && invalid() ? styles.Invalid : ''}`}>
                        {board()[props.id]}
                    </h2>
                )
            }
        </div>
    );
};