// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

using Formulator.Core.Entities;

namespace Formulator.Infrastructure.Data
{
    internal class ApplicationDbInitialiser
    {
        public static void Initialise(ApplicationDbContext context)
        {
            // todo: instead ensure migrations
            context.Database.EnsureCreated();

            // Check if the database has already been seeded
            if (context.Users.Any())
            {
                return;
            }

            // Add seed data
            context.AddRange(new ApplicationUser[]
            {
                new ApplicationUser{
                    Id = 1,
                    DisplayName = "user1",
                    Email = "user1@localhost"
                },
                new ApplicationUser{
                    Id = 2,
                    DisplayName = "user2",
                    Email = "user2@localhost"
                },
            });

            context.SaveChanges();
        }
    }
}
