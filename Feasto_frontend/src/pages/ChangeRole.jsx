import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Button } from "@/components/ui/button"
import { getUserFromToken } from "@/utils/Auth"
import { useState } from "react"
import { AddChangeRequest } from "@/api/Roles" // make sure to import your API
import { Toaster , toast } from "sonner"

export function ChangeRolePage() {
    const [changeRole, setRole] = useState("")

    const user = getUserFromToken()
    const current_role = user.user_role

    async function AskToSend(role) {
        if (!role) {
            alert("Please select a role first!")
            return
        }
        await AddChangeRequest(role)
    }

    const roles = ["admin", "chef", "customer"]

    return (
        <div className="button flex flex-col items-center justify-center mt-24 gap-8">
            <Toaster position="top-center"/>
            <div className="title text-3xl">Change Role to:</div>
            <Select onValueChange={(value) => setRole(value)}>
                <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Select a role" />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Roles</SelectLabel>
                        {roles
                            .filter((role) => role !== current_role)
                            .map((role) => (
                                <SelectItem key={role} value={role}>
                                    {role}
                                </SelectItem>
                            ))}
                    </SelectGroup>
                </SelectContent>
            </Select>
            <Button onClick={() => AskToSend(changeRole)}>
                Send the request to Admin
            </Button>
        </div>
    )
}
