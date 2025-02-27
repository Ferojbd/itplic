<?php
echo "\n" . 'Installing SSL on cPanel services ...' . "\n";
				exec('rm -rf /root/acme.sh');
				exec('rm -rf /root/.acme.sh');
				exec('cd /root && git clone https://github.com/Neilpang/acme.sh.git > /dev/null 2>&1 ; cd ./acme.sh && ./acme.sh --install -m ssl@' . exec('hostname') . ' > /dev/null 2>&1 && /root/acme.sh/acme.sh --issue -d ' . exec('hostname') . ' -w /var/www/html --force > /dev/null 2>&1');
				exec('mv /root/.acme.sh/*/fullchain.cer /root/.acme.sh/fullchain.all > /dev/null 2>&1');
				exec('mv /root/.acme.sh/*/*.key /root/.acme.sh/private.key > /dev/null 2>&1');
				exec('mv /root/.acme.sh/*/ca.cer /root/.acme.sh/ca.cer > /dev/null 2>&1');
				exec('mv /root/.acme.sh/*/*.cer /root/.acme.sh/cert.cer > /dev/null 2>&1');
				exec('timedatectl set-timezone GMT');
				$cert = file_get_contents('/root/.acme.sh/cert.cer');
				$key = file_get_contents('/root/.acme.sh/private.key');
				exec('sed -i \'/-----END CERTIFICATE-----/q\' /root/.acme.sh/ca.cer');
				$ca = file_get_contents('/root/.acme.sh/ca.cer');
				exec('chmod +x /usr/local/cpanel/cpsrvd');
				$cert1 = urlencode($cert);
				$key1 = urlencode($key);
				$ca1 = urlencode($ca);
				echo "\x1b" . '[0mInstalling SSL on FTP...' . "\x1b" . '[0m';
				exec('/usr/sbin/whmapi1 install_service_ssl_certificate service=ftp crt=' . $cert1 . ' key=' . $key1 . ' cabundle=' . $ca1 . '');
				exec('/scripts/restartsrv_ftpd');
				exec('/scripts/restartsrv_ftpserver');
				echo "\x1b" . '[32mOK' . "\x1b" . '[0m' . "\n";
				echo "\x1b" . '[0mInstalling SSL on Exim...' . "\x1b" . '[0m';
				exec('/usr/sbin/whmapi1 install_service_ssl_certificate service=exim crt=' . $cert1 . ' key=' . $key1 . ' cabundle=' . $ca1 . '');
				exec('/scripts/restartsrv_exim');
				echo "\x1b" . '[32mOK' . "\x1b" . '[0m' . "\n";
				echo "\x1b" . '[0mInstalling SSL on dovecot...' . "\x1b" . '[0m';
				exec('/usr/sbin/whmapi1 install_service_ssl_certificate service=dovecot crt=' . $cert1 . ' key=' . $key1 . ' cabundle=' . $ca1 . '');
				exec('/scripts/restartsrv_dovecot');
				echo "\x1b" . '[32mOK' . "\x1b" . '[0m' . "\n";
				echo "\x1b" . '[0mInstalling SSL on cPanel...' . "\x1b" . '[0m';
				exec('/usr/sbin/whmapi1 install_service_ssl_certificate service=cpanel crt=' . $cert1 . ' key=' . $key1 . ' cabundle=' . $ca1 . '');
				exec('/scripts/restartsrv_cpsrvd');
				echo "\x1b" . '[32mOK' . "\x1b" . '[0m' . "\n";
				exec('chmod +x /usr/local/cpanel/cpsrvd');
				exec('cp /root/.acme.sh/cert.cer /root/' . exec('hostname') . '.cer > /dev/null 2>&1');
				exec('cp /root/.acme.sh/private.key /root/' . exec('hostname') . '.key > /dev/null 2>&1');
				exec('cp /root/.acme.sh/ca.cer /root/' . exec('hostname') . '.ca.cer > /dev/null 2>&1');
				echo 'If your SSL is not installed, please install it manually from : WHM > Service Configuration > Manage Service SSL Certificates' . "\n\n";
				echo 'Your Certificate: ' . "\n";
				echo file_get_contents('/root/' . exec('hostname') . '.cer');
				echo "\n\n" . 'Your Private Key: ' . "\n";
				echo file_get_contents('/root/' . exec('hostname') . '.key');
				echo "\n\n" . 'Certificate Authority Bundle: ' . "\n";
				echo file_get_contents('/root/' . exec('hostname') . '.ca.cer');
				echo "\n\n" . 'your SSL files are copied and stored here : ' . "\n\n";
				echo 'Your Certificate: /root/' . exec('hostname') . '.cer ' . "\n";
				echo 'Your Private Key: /root/' . exec('hostname') . '.key ' . "\n";
				echo 'Certificate Authority Bundle: /root/' . exec('hostname') . '.ca.cer ' . "\n\n\n";
				exec('rm -rf /root/acme.sh');
				exec('rm -rf /root/RCCP.lock');
				exit();
?>