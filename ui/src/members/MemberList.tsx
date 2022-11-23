import React, { useState } from 'react';
import { Member } from './Member';
import MemberListRow from './MemberListRow';
import MemberForm from './MemberForm';
import MemberListRowHeader from './MemberListRowHeader';

interface MemberListProps {
    members: Member[];
    onSave: (member: Member) => void;
}

function MemberList(props: MemberListProps) {
    const { members, onSave } = props;
    const [memberBeingEdited, setMemberBeingEdited] = useState({});
    const handleEdit = (member: Member) => {
        // console.log(member);
        setMemberBeingEdited(member);
    };
    const cancelEditing = () => {
        setMemberBeingEdited({});
    };
    return (
        <div className="container">
            <div className="row"><MemberListRowHeader /></div>
            {members.map((member) => (
                <div>
                    {member === memberBeingEdited ? (
                        <MemberForm member={member} onSave={onSave} onCancel={cancelEditing}/>
                    ) : (
                        <MemberListRow member={member} onEdit={handleEdit}/>
                    )}
                </div>
            ))}
        </div>
    );
}

export default MemberList;