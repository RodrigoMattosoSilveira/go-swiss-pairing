import React from "react";
import {Member} from "./Member";


interface MemberFormModalProps {
    member: Member;
}

function MemberFormModal(props: MemberFormModalProps) {
    const { member: member } = props;
    return (
        <>
            <label htmlFor="modal-control">Edit</label>
            <input type="checkbox" id="modal-control" className="modal"></input>
            <div role="dialog" aria-labelledby="dialog-title">
                <div className="card">
                    <label htmlFor="modal-control" className="modal-close"></label>
                    <h3 className="section" id="dialog-title">Modal</h3>
                    <p className="section">This is a modal dialog! {member.id}</p>
                </div>
            </div>
        </>
    );
}

export default MemberFormModal;