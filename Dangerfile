message "Welcome to danger"

# # Sometimes it's a README fix, or something like that - which isn't relevant for
# # including in a project's CHANGELOG for example
# declared_trivial = github.pr_title.include? "#trivial"

# # Make it more obvious that a PR is a work in progress and shouldn't be merged yet
# warn("PR is classed as Work in Progress") if github.pr_title.include? "[WIP]"

# # Warn when there is a big PR
# warn("Big PR") if git.lines_of_code > 500

# # Don't let testing shortcuts get into master by accident
# fail("fdescribe left in tests") if `grep -r fdescribe specs/ `.length > 1
# fail("fit left in tests") if `grep -r fit specs/ `.length > 1

newNoEncryptOptions = 0
git.diff.each do |chunk|
    if chunk.path.match(/src\/el/)
      chunk.patch.lines.grep(/^+/).each do |added_line|
          if added_line.match(/NoEncrypt:\s*true/)
            newNoEncryptOptions = newNoEncryptOptions + 1
          end
      end
      chunk.patch.lines.grep(/^-/).each do |removed_line|
          if removed_line.match(/NoEncrypt:\s*true/)
            newNoEncryptOptions = newNoEncryptOptions - 1
          end
      end
    end
end
message "Result: #{newNoEncryptOptions}"
# if newNoEncryptOptions > 0
#   begin
#     github.api.request_pull_request_review('EverlongProject/services', github.pr_json["number"], reviewers: ['@EverlongProject/productivity-backend'])
#   rescue StandardError => e
#     message "Error requesting backend-productivity review: #{e.message}"
#     raise
#   end
# end