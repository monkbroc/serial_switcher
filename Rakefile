TARGET="serial_switcher"
TARGETS = [
  {os: 'windows', arch: '386'},
  {os: 'windows', arch: 'amd64'},
  {os: 'darwin', arch: 'amd64'},
  {os: 'linux', arch: 'arm', goarm: '6'},
  {os: 'linux', arch: 'amd64'},
  {os: 'linux', arch: '386', go386: '387'},
]

desc "build particle-cli-ng"
task :build do
  puts "Building..."
  FileUtils.mkdir_p 'dist'
  TARGETS.each do |target|
    puts "  * #{target[:os]}-#{target[:arch]}"
    build(target)
  end
end

def build(target)
  path = "./dist/#{target[:os]}/#{target[:arch]}/#{TARGET}"
  path += ".exe" if target[:os] === 'windows'
  # ldflags = "-X=main.Version=#{VERSION} -X=main.Channel=#{CHANNEL}"
  args = ["-o", "#{path}"]
  # args += ["-ldflags", "\"#{ldflags}\""]
  # unless target[:os] === 'windows'
  #   args += ["-a", "-tags", "netgo"]
  # end
  vars = ["GOOS=#{target[:os]}", "GOARCH=#{target[:arch]}"]
  vars << "GO386=#{target[:go386]}" if target[:go386]
  vars << "GOARM=#{target[:goarm]}" if target[:goarm]
  ok = system("#{vars.join(' ')} go build #{args.join(' ')}")
  exit 1 unless ok
end

task :default => :build
