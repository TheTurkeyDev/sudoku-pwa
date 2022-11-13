import { JSXElement, Show } from 'solid-js';
import styles from './modal.module.css';

type ModalProps = {
    readonly children: JSXElement
    readonly show: boolean
}

export const Modal = (props: ModalProps) => {
    return (
        <Show when={props.show}>
            <div class={styles.BackgroundWrapper}>
                <div class={styles.ModalContent}>
                    {props.children}
                </div>
            </div>
        </Show>
    );
};