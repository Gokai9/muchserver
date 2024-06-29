using Microsoft.EntityFrameworkCore;

namespace animecsharpapilist.Models;

public class AnimeContext: DbContext {
    public AnimeContext(DbContextOptions<AnimeContext> options): base(options) {
        
    }
    public DbSet<AnimeList> AnimeLists {get; set;}
}