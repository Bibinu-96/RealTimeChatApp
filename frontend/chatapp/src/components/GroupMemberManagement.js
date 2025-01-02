import React, { useState } from 'react';
import { Search, UserPlus, X, Mail, User } from 'lucide-react';
import { 
  Card,
  CardHeader,
  CardTitle,
  CardContent,
} from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Alert,
  AlertDescription,
} from "@/components/ui/alert";

const GroupMemberManagement = () => {
  const [members, setMembers] = useState([
    { id: 1, name: 'John Doe', email: 'john@example.com', avatar: '/api/placeholder/32/32' },
    { id: 2, name: 'Jane Smith', email: 'jane@example.com', avatar: '/api/placeholder/32/32' }
  ]);
  
  const [searchTerm, setSearchTerm] = useState('');
  const [newMemberName, setNewMemberName] = useState('');
  const [newMemberEmail, setNewMemberEmail] = useState('');
  const [showSuccessAlert, setShowSuccessAlert] = useState(false);
  const [suggestions] = useState([
    { id: 3, name: 'Alice Johnson', email: 'alice@example.com', avatar: '/api/placeholder/32/32' },
    { id: 4, name: 'Bob Wilson', email: 'bob@example.com', avatar: '/api/placeholder/32/32' },
    { id: 5, name: 'Carol Brown', email: 'carol@example.com', avatar: '/api/placeholder/32/32' }
  ]);

  const addMember = (newMember) => {
    if (!members.find(member => member.id === newMember.id)) {
      setMembers([...members, newMember]);
    }
  };

  const handleAddNewMember = () => {
    if (newMemberName && newMemberEmail) {
      const newMember = {
        id: Date.now(), // temporary ID generation
        name: newMemberName,
        email: newMemberEmail,
        avatar: '/api/placeholder/32/32'
      };
      
      addMember(newMember);
      setNewMemberName('');
      setNewMemberEmail('');
      setShowSuccessAlert(true);
      setTimeout(() => setShowSuccessAlert(false), 3000);
    }
  };

  const removeMember = (memberId) => {
    setMembers(members.filter(member => member.id !== memberId));
  };

  const filteredSuggestions = suggestions.filter(user => 
    !members.find(member => member.id === user.id) &&
    (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
     user.email.toLowerCase().includes(searchTerm.toLowerCase()))
  );

  return (
    <Card className="w-full max-w-md">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle>Group Members</CardTitle>
        <Dialog>
          <DialogTrigger className="inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
            <UserPlus className="h-4 w-4 mr-2" />
            Add Member
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Add New Member</DialogTitle>
            </DialogHeader>
            <div className="space-y-4 py-4">
              <div className="space-y-2">
                <label className="text-sm font-medium">Name</label>
                <div className="relative">
                  <User className="absolute left-2 top-2.5 h-4 w-4 text-gray-400" />
                  <input
                    type="text"
                    placeholder="Enter member name"
                    className="pl-8 pr-4 py-2 w-full border rounded-md"
                    value={newMemberName}
                    onChange={(e) => setNewMemberName(e.target.value)}
                  />
                </div>
              </div>
              <div className="space-y-2">
                <label className="text-sm font-medium">Email</label>
                <div className="relative">
                  <Mail className="absolute left-2 top-2.5 h-4 w-4 text-gray-400" />
                  <input
                    type="email"
                    placeholder="Enter member email"
                    className="pl-8 pr-4 py-2 w-full border rounded-md"
                    value={newMemberEmail}
                    onChange={(e) => setNewMemberEmail(e.target.value)}
                  />
                </div>
              </div>
              <button
                onClick={handleAddNewMember}
                className="w-full inline-flex justify-center items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                disabled={!newMemberName || !newMemberEmail}
              >
                Add Member
              </button>
            </div>
          </DialogContent>
        </Dialog>
      </CardHeader>
      <CardContent>
        {showSuccessAlert && (
          <Alert className="mb-4 bg-green-50 border-green-200">
            <AlertDescription className="text-green-800">
              Member added successfully!
            </AlertDescription>
          </Alert>
        )}

        {/* Search Bar */}
        <div className="relative mb-4">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-gray-400" />
          <input
            type="text"
            placeholder="Search users..."
            className="pl-8 pr-4 py-2 w-full border rounded-md"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>

        {/* Member List */}
        <div className="space-y-2 mb-4">
          <h3 className="text-sm font-medium">Current Members ({members.length})</h3>
          {members.map(member => (
            <div key={member.id} className="flex items-center justify-between p-2 bg-gray-50 rounded-md">
              <div className="flex items-center space-x-2">
                <img
                  src={member.avatar}
                  alt={member.name}
                  className="w-8 h-8 rounded-full"
                />
                <div>
                  <div className="font-medium">{member.name}</div>
                  <div className="text-sm text-gray-500">{member.email}</div>
                </div>
              </div>
              <button
                onClick={() => removeMember(member.id)}
                className="text-gray-400 hover:text-gray-600"
              >
                <X className="h-4 w-4" />
              </button>
            </div>
          ))}
        </div>

        {/* Suggestions */}
        {searchTerm && filteredSuggestions.length > 0 && (
          <div className="space-y-2">
            <h3 className="text-sm font-medium">Suggestions</h3>
            {filteredSuggestions.map(user => (
              <div key={user.id} className="flex items-center justify-between p-2 bg-gray-50 rounded-md">
                <div className="flex items-center space-x-2">
                  <img
                    src={user.avatar}
                    alt={user.name}
                    className="w-8 h-8 rounded-full"
                  />
                  <div>
                    <div className="font-medium">{user.name}</div>
                    <div className="text-sm text-gray-500">{user.email}</div>
                  </div>
                </div>
                <button
                  onClick={() => addMember(user)}
                  className="text-blue-500 hover:text-blue-600"
                >
                  <UserPlus className="h-4 w-4" />
                </button>
              </div>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  );
};

export default GroupMemberManagement;