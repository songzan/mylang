#! perl

use 5.012;
use experimental 'switch';

$|=1;

use constant {
  FOF   => chr(0),
  EIN   => chr(1),
  OUT   => chr(2),
  QSTR  => chr(3),
  FAIL  => chr(0),
  PASS  => chr(1),
};

sub error {
  my $message = shift;
  say $message;
  exit;
}

sub now { return time(); }

sub readlin { return <STDIN> }

sub startwith {
  my ($str,$start) = @_;
  return 1 if index($str, $start) == 0;
  return 0;
}

sub endwith {
  my ($s, $c) = @_;
  return 1 if tail($s) eq $c;
  return 0;
}

sub first {
  my $data = shift;
  return substr($data,0,1);
}

sub second {
  my $str = shift;
  return substr($str,1,1);
}

sub tail {
  my $str = shift;
  return substr($str,-1);
}

sub indexat {
  my ($str, $at) = @_;
  return substr($str,$at,1);
}

sub cut {
  my ($s) = @_;
  return substr($s,1,-1);
}

sub chopstr {
  my ($s) = @_;
  return substr($s,0,-1);
}

sub rest {
  my ($s) = @_;
  return substr($s,1);
}

sub restat {
  my ($s,$at) = @_;
  return substr($s,$at);
}

sub toend {
  my ($s) = @_;
  if (length($s) < 40) { return $s; }
  return substr($s,0,40);
}

sub chars {
  my $str = shift;
  my $len = length($str) - 1;
  my @chars;
  for my $i (0 .. $len) {
    push @chars, substr($str,$i,1)
  }
  return @chars;
}

sub hkeys {
  my $hash = shift;
  return keys(%{$hash});
}

sub tochars {
  my $str = shift;
  return [chars($str)];
}

sub splitstr {
  my ($str, $char) = @_;
  return [ split quotemeta($char), $str ];
}

sub itos {
  my $int = shift;
  return "$int";
}

sub len { 
  my $arr = shift;
  return scalar(@{$arr});
}

sub string {
  my $arr = shift;
  return join '', @{$arr};
}

sub joinstrs {
  my ($arr, $char) = @_;
  return join($char, @{$arr});
}

sub reststrs {
  my $data = shift;
  my @arr = @{$data};
  return [splice @arr, 1];
}

sub include {
  my ($s, $bs) = @_;
  return 1 if index($s,$bs) >= 0;
  return 0;
}

sub has {
  my ($hash, $key) = @_;
  return 1 if exists $hash->{$key};
  return 0;
}

sub readfile {
  my $file = shift;
  error("$file not exists") if not -e $file;
  local $/;
  open my ($fh), '<', $file or die $!;
  return <$fh>;
}

sub writefile {
  my ($file,$str) = @_;
  open my ($fh), '>', $file or die $!;
  print {$fh} $str;
  return $file;
}

sub toint {
  my $str = shift;
  return 0 + $str;
}

sub osargs { return [@ARGV]; }

sub getline {
  my ($str, $off) = @_;
  my $s = substr($str,0,$off);
  my $count = 0;
  for my $ch (chars($s)) {
    $count++ if $ch eq "\n";
  }
  return $count
}

sub trim {
  my $str = shift;
  $str =~ s/^\s+|\s+$//g;
  return $str;
}

sub repeat {
  my ($a,$b) = @_;
  return ($a x $b);
}

sub isfile {
  my $file = shift;
  if (-e $file) { return 1; }
  return 0;
}

sub isupper {
  my ($c) = @_;
  if ($c ge 'A') {
    if ($c le 'Z') { return 1; }
  }
  return 0;
}

sub isvspace {
  my ($c) = @_;
  if ($c eq "\n") { return 1; }
  if ($c eq "\r") { return 1; }
  return 0;
}

sub ishspace {
  my ($c) = @_;
  if ($c eq ' ') { return 1; }
  if ($c eq "\t") { return 1; }
  return 0;
}

sub isspace {
  my ($c) = @_;
  if (isvspace($c)) { return 1; }
  if (ishspace($c)) { return 1; }
  return 0;
}

sub isdigit {
  my ($c) = @_;
  if ($c lt '0') { return 0; }
  if ($c gt '9') { return 0; }
  return 1;
}

sub islower {
  my ($c) = @_;
  if ($c ge 'a') {
    if ($c le 'z') { return 1; }
  }
  return 0;
}

sub isletter {
  my ($c) = @_;
  if (islower($c)) { return 1; }
  if (isupper($c)) { return 1; }
  return 0;
}

sub isalpha {
  my ($c) = @_;
  if (isletter($c)) { return 1; }
  if ($c eq '_') { return 1; }
  return 0;
}

sub iswords {
  my ($c) = @_;
  if (isalpha($c)) { return 1; }
  if (isdigit($c)) { return 1; }
  if ($c eq '-') { return 1; }
  return 0;
}

sub isxdigit {
  my ($c) = @_;
  if (isdigit($c)) { return 1; }
  if ($c ge 'A') {
    if ($c le 'F') { return 1; }
  }
  if ($c ge 'a') {
    if ($c le 'f') { return 1; }
  }
  return 0;
}
