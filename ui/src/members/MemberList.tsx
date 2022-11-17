import React from 'react';
import { Member } from './Member';

interface ProjectListProps {
    members: Member[];
}

function ProjectList({ members }: ProjectListProps) {
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
   return (
         <div className="row">
               {members.map((member) => (
                 <div key={member.id} className="cols-sm">
                       <div className="card">
                         <img src={"assets/" + member.imageUrl} alt={member.first + " " + member.last} />
                         <section className="section dark">
                           <h5 className="strong">
                             <strong>{member.first + " " + member.last}</strong>
                           </h5>
                           <p>Cell: {member.cell}</p>
                           <p>Email {member.email}</p>
                           <p>Rating : {member.rating.toLocaleString()}</p>
                         </section>
                       </div>
                     </div>
               ))}
             </div>
       );
}

export default ProjectList;