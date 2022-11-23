import React from 'react';
import { Link } from 'react-router-dom';
import { Member } from './Member';
import ErrorBoundary from "../util/ErrorBoundary";

interface MemberListRowProps {
    member: Member;
    onEdit: (member: Member) => void;
}

function MemberListRow(props: MemberListRowProps) {
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
    const { member, onEdit } = props;
    const handleEditClick = (memberBeingEdited: Member) => {
        // console.log(memberBeingEdited);
        onEdit(memberBeingEdited);
    };
    return (
        <div key={member.id} className="row">
            <div className="col-sm-1">
                <button
                    className="small"
                    aria-label={"member.id"}
                    onClick={() => {
                        handleEditClick(member);
                    }}
                >
                    <span className="icon-edit"></span>
                </button>
            </div>
            <ErrorBoundary><Link to={'/members/' + member.id} className="col-sm-1">{member.id}</Link></ErrorBoundary>
            <div className="col-sm-1" aria-label={"member.first"}>{member.first}</div>
            <div className="col-sm-1" aria-label={"member.last"}>{member.last}</div>
            <div className="col-sm-2" aria-label={"member.email"}>{member.email}</div>
            <div className="col-sm-2" aria-label={"member.cell"}>{member.cell}</div>
            <div className="col-sm-2" aria-label={"member.password"}>{member.password}</div>
            <div className="col-sm-1" aria-label={"member.rating"}>{member.rating}</div>
            <div className="col-sm-1" aria-label={"member.isActive"}>{member.isActive ? "Yes" : "No"}</div>
        </div>
    )
}

export default MemberListRow;