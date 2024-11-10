import { useState } from 'react';
import { Box, Button } from '@mui/joy';
import WorkTime from '@/features/work-time/components/work-time';
import Timer from '@/features/timer/components/timer';

export default function SidebarCardDown() {
  const [activeTab, setActiveTab] = useState('work');

  const handleTabChange = (tab: string) => {
    setActiveTab(tab);
  };

  return (
    <Box
      sx={{
        p: 2,
        borderRadius: '8px',
        backgroundColor: '#f3f0e9',
        mb: 3,
      }}
    >
      <Box display="flex" mb={2} borderRadius="8px" overflow="hidden">
        <Button
          onClick={() => handleTabChange('work')}
          variant={activeTab === 'work' ? 'solid' : 'plain'}
          color={activeTab === 'work' ? 'success' : 'neutral'}
          sx={{ flex: 1, fontWeight: 'bold', borderRadius: 0 }}
        >
          作業時間
        </Button>
        <Button
          onClick={() => handleTabChange('timer')}
          variant={activeTab === 'timer' ? 'solid' : 'plain'}
          color={activeTab === 'timer' ? 'success' : 'neutral'}
          sx={{ flex: 1, fontWeight: 'bold', borderRadius: 0 }}
        >
          タイマー
        </Button>
      </Box>
      {activeTab === 'work' ? <WorkTime /> : <Timer />}
    </Box>
  );
}
