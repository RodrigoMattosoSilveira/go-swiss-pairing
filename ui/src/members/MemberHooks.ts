import { useState, useEffect } from 'react';
import { memberAPI } from './memberAPI';
import { Member } from './Member';

export function useMembers() {
    const [members, setMembers] = useState<Member[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | undefined>(undefined);
    const [currentPage, setCurrentPage] = useState(1);
    const [saving, setSaving] = useState(false);
    const [savingError, setSavingError] = useState<string | undefined>(
        undefined
    );

    useEffect(() => {
        async function loadMembers() {
            setMembers([]);
            setLoading(true);
            try {
                const data = await memberAPI.get(currentPage);
                if (currentPage === 1) {
                    setMembers(data);
                } else {
                    setMembers((members) => [...members, ...data]);
                }
            } catch (e) {
                if (e instanceof Error) {
                    setError(e.message);
                }
            } finally {
                setLoading(false);
            }
        }
        loadMembers();
    }, [currentPage]);

    const saveMember = (member: Member) => {
        setSaving(true);
        memberAPI
            .put(member)
            .then((updatedMember) => {
                let updatedMembers = members.map((p) => {
                    return p.id === member.id ? new Member(updatedMember) : p;
                });
                setMembers(updatedMembers);
            })
            .catch((e) => {
                setSavingError(e.message);
            })
            .finally(() => setSaving(false));
    };

    return {
        members,
        loading,
        error,
        currentPage,
        setCurrentPage,
        saving,
        savingError,
        saveMember,
    };
}