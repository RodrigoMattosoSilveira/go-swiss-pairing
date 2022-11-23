import React, { useEffect, useState } from 'react';
import { memberAPI } from './memberAPI';
import MemberDetail from './MemberDetail';
import { Member } from './Member';
import { useParams } from 'react-router-dom';

function MemberPage(props: any) {
    const [loading, setLoading] = useState(false);
    const [member, setMember] = useState<Member | null>(null);
    const [error, setError] = useState<string | null>(null);
    const params = useParams();
    const id = params.id;

    useEffect(() => {
        setLoading(true);
        if (id != null) {
            memberAPI
                .find(id)
                .then((data) => {
                    // @ts-ignore
                    setMember(data);
                    setLoading(false);
                })
                .catch((e) => {
                    setError(e);
                    setLoading(false);
                });
        }
    }, [id]);

    return (
        <div>
            <>
                <h1>Member Detail</h1>

                {loading && (
                    <div className="center-page">
                        <span className="spinner primary"></span>
                        <p>Loading...</p>
                    </div>
                )}

                {error && (
                    <div className="row">
                        <div className="card large error">
                            <section>
                                <p>
                                    <span className="icon-alert inverse "></span> {error}
                                </p>
                            </section>
                        </div>
                    </div>
                )}
                {member && <MemberDetail member={member} />}
            </>
        </div>
    );
}

export default MemberPage;