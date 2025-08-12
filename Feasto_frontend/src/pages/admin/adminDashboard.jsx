import { Navbar } from "@/components/Navbar"

export function AdminDashboard(){
    return (
        <>
        <Navbar user="admin" page="AdminDashboard"/>
        <div className="content">
            <div className="flex justify-center font-extrabold title text-5xl text-green-500 mt-16">
                Welcome Admin!!!!!
            </div>
        </div>
        </>
    )
}