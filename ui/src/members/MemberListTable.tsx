import React from 'react';
import { Member } from './Member';
import MemberListTableRow from "./MemberListTableRow";

interface MemberListTableProps {
    members: Member[];
}

function MemberListTable({ members }: MemberListTableProps) {
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
   return (
         // <div className="row">
         //       {members.map((member) => (
         //         <div key={member.id} className="cols-sm">
         //               <div className="card">
         //                 <img src={"assets/" + member.imageUrl} alt={member.first + " " + member.last} />
         //                 <section className="section dark">
         //                   <h5 className="strong">
         //                     <strong>{member.first + " " + member.last}</strong>
         //                   </h5>
         //                   <p>Cell: {member.cell}</p>
         //                   <p>Email {member.email}</p>
         //                   <p>Rating : {member.rating.toLocaleString()}</p>
         //                 </section>
         //               </div>
         //             </div>
         //       ))}
         //     </div>
        <table className="hoverable">
            {/*<caption>Members</caption>*/}
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Cell</th>
                    <th>Rating</th>
                    <th>Active</th>
                    <th>Edit</th>
                </tr>
            </thead>
            <tbody>
                {members.map((member) => (
                    <tr key={member.id}>
                        <MemberListTableRow member={member}></MemberListTableRow>
                    </tr>
                ))}
            </tbody>
        </table>
   );
}

export default MemberListTable;