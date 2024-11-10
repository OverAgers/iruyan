import SidebarCardDown from '@/components/ui/card/sidebar-card-down';
import SidebarCardUp from '@/components/ui/card/sidebar-card-up';
import { Box, Typography } from '@mui/joy';
import useCurrentTime from '@/hooks/current-time-hooks';

export default function Sidebar() {
  const currentTime = useCurrentTime();
  return (
    <Box justifySelf={'flex-end'} pr={2} pt={3}>
      <SidebarCardUp />
      <SidebarCardDown />
      <Box>
        <Typography level="h2" mb={2} sx={{ color: '#F7F4ED' }}>
          現在時刻
        </Typography>
        <Typography fontSize={'128px'} lineHeight={1} sx={{ color: '#F7F4ED' }}>
          {currentTime
            ? currentTime.toLocaleTimeString([], {
                hour: '2-digit',
                minute: '2-digit',
              })
            : '--:--'}
        </Typography>
      </Box>
    </Box>
  );
}
