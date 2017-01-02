require 'date'

f = File.open('./test.txt', 'r')
f.each_line do |line|
    dl                         = line[0..9]
    last_name       = line[10..49].strip
    first_name         = line[50..89]&.strip
    middle_name     = line[90..128]&.strip
    suffix                 = line[129..134]&.strip
    dob                     = line[135..142]&.strip
    dob                            = Date.strptime(dob, "%m%d%Y") if dob
    address1             = line[143..174]&.strip&.gsub(/\s{2,}/, ' ')
    address2             = line[175..206]&.strip&.gsub(/\s{2,}/, ' ')
    city                     = line[207..239]&.strip
    state                 = line[240..241]&.strip
    zip_code             = line[242..246]&.strip
    zip_code_ext         = line[247..250]&.strip
    orig_issue_date = line[251..258]&.strip
    orig_issue_date = Date.strptime(orig_issue_date, "%m%d%Y") if orig_issue_date
    card_type             = line[259..261]&.strip
    puts "#{dl},#{first_name},#{middle_name},#{last_name},#{suffix},#{dob},#{address1},#{address2},#{city},#{state},#{zip_code},#{zip_code_ext},#{orig_issue_date},#{card_type}"
end

f.close